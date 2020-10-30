package relay

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"

	"marlin-cli/util"
)


func CreateCommand() *cli.Command {
	var discovery_addrs, heartbeat_addrs, datadir string
	var discovery_port, pubsub_port uint
	var address, name string

	return &cli.Command{
		Name:  "create",
		Usage: "create a new relay",
		Flags: []cli.Flag{
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
			out, _ := exec.Command("sudo", "supervisorctl", "status", "relay").Output()
			if !strings.Contains(string(out), "no such process") {
				// Already exists
				return errors.New("Already exists")
			}

			// User details
			usr, err := util.GetUser()
			if err != nil {
				return err
			}

			// relay executable
			err = util.Fetch("https://storage.googleapis.com/marlin-artifacts/bin/relay", usr.HomeDir+"/.marlin/ctl/bin/relay", usr.Username, true)
			if err != nil {
				return err
			}

			// relay config
			err = util.Fetch("https://storage.googleapis.com/marlin-artifacts/configs/relay.conf", usr.HomeDir+"/.marlin/ctl/configs/relay.conf", usr.Username, false)
			if err != nil {
				return err
			}

			err = util.TemplatePlace(
				usr.HomeDir+"/.marlin/ctl/configs/relay.conf",
				"/etc/supervisor/conf.d/relay.conf",
				struct {
					Program, User, UserHome string
					DiscoveryAddrs, HeartbeatAddrs, Datadir string
					DiscoveryPort, PubsubPort uint
					Address, Name string
				} {
					"relay", usr.Username, usr.HomeDir,
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

			_, err = exec.Command("supervisorctl", "add", "relay").Output()
			if err != nil {
				return err
			}

			return nil
		},
	}
}
