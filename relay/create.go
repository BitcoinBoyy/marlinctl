package relay

import (
	"errors"
	"os/exec"
	"strings"
	"runtime"

	"github.com/urfave/cli/v2"

	"marlin-cli/util"
)


func CreateCommand() *cli.Command {
	var chain, discovery_addrs, heartbeat_addrs, datadir string
	var discovery_port, pubsub_port uint
	var address, name string

	return &cli.Command{
		Name:  "create",
		Usage: "create a new relay",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "chain",
				Usage:       "--chain \"<CHAIN>\"",
				Destination: &chain,
				Required: true,
			},
			&cli.StringFlag{
				Name:        "discovery-addrs",
				Usage:       "--discovery-addrs \"<IP1:PORT1>,<IP2:PORT2>,...\"",
				Destination: &discovery_addrs,
				Required: true,
			},
			&cli.StringFlag{
				Name:        "heartbeat-addrs",
				Usage:       "--heartbeat-addrs \"<IP1:PORT1>,<IP2:PORT2>,...\"",
				Destination: &heartbeat_addrs,
				Required: true,
			},
			&cli.StringFlag{
				Name:        "datadir",
				Usage:       "--datadir \"/path/to/datadir\"",
				Destination: &datadir,
				Required: true,
			},
			&cli.UintFlag{
				Name:        "discovery-port",
				Usage:       "--discovery-port <PORT>",
				Destination: &discovery_port,
			},
			&cli.UintFlag{
				Name:        "pubsub-port",
				Usage:       "--pubsub-port <PORT>",
				Destination: &pubsub_port,
			},
			&cli.StringFlag{
				Name:        "address",
				Usage:       "--address \"0x...\"",
				Destination: &address,
			},
			&cli.StringFlag{
				Name:        "name",
				Usage:       "--name \"<NAME>\"",
				Destination: &name,
			},
		},
		Action: func(c *cli.Context) error {
			program := chain+"_relay"

			out, _ := exec.Command("sudo", "supervisorctl", "status", program).Output()
			if !strings.Contains(string(out), "no such process") {
				// Already exists
				return errors.New("Already exists")
			}

			// Set up abci first
			if abci, found := abciMap[chain]; found {
				err := abci.Create(datadir)
				if err != nil {
					return err
				}
			} else {
				return errors.New("Unrecognized chain")
			}

			// User details
			usr, err := util.GetUser()
			if err != nil {
				return err
			}

			// relay executable
			err = util.Fetch("https://storage.googleapis.com/marlin-artifacts/bin/"+program+"-"+runtime.GOOS+"-"+runtime.GOARCH, usr.HomeDir+"/.marlin/ctl/bin/"+program, usr.Username, true, false)
			if err != nil {
				return err
			}

			// relay config
			err = util.Fetch("https://storage.googleapis.com/marlin-artifacts/configs/"+program+".conf", usr.HomeDir+"/.marlin/ctl/configs/"+program+".conf", usr.Username, false, false)
			if err != nil {
				return err
			}

			err = util.TemplatePlace(
				usr.HomeDir+"/.marlin/ctl/configs/"+program+".conf",
				"/etc/supervisor/conf.d/"+program+".conf",
				struct {
					Program, User, UserHome string
					DiscoveryAddrs, HeartbeatAddrs, Datadir string
					DiscoveryPort, PubsubPort uint
					Address, Name string
				} {
					program, usr.Username, usr.HomeDir,
					discovery_addrs, heartbeat_addrs, datadir,
					discovery_port, pubsub_port,
					address, name,
				},
			)
			if err != nil {
				return err
			}

			_, err = exec.Command("supervisorctl", "reread").Output()
			if err != nil {
				return err
			}

			_, err = exec.Command("supervisorctl", "add", program).Output()
			if err != nil {
				return err
			}

			return nil
		},
	}
}
