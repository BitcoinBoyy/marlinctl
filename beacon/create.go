package beacon

import (
	"errors"
	"os/exec"
	"strings"
	"runtime"

	"github.com/urfave/cli/v2"

	"marlin-cli/util"
)


func CreateCommand() *cli.Command {
	var discovery_addr string
	var heartbeat_addr string
	var beacon_addr string
	var program string

	return &cli.Command{
		Name:  "create",
		Usage: "create a new beacon",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "program",
				Usage:       "--program <NAME>",
				Value:       "beacon",
				Destination: &program,
			},
			&cli.StringFlag{
				Name:        "discovery-addr",
				Usage:       "--discovery-addr <IP:PORT>",
				Destination: &discovery_addr,
			},
			&cli.StringFlag{
				Name:        "heartbeat-addr",
				Usage:       "--heartbeat-addr <IP:PORT>",
				Destination: &heartbeat_addr,
			},
			&cli.StringFlag{
				Name:        "beacon-addr",
				Usage:       "--beacon-addr <IP:PORT>",
				Destination: &beacon_addr,
			},
		},
		Action: func(c *cli.Context) error {
			out, _ := exec.Command("sudo", "supervisorctl", "status", program).Output()
			if !strings.Contains(string(out), "no such process") {
				// Already exists
				return errors.New("Already exists")
			}

			// User details
			usr, err := util.GetUser()
			if err != nil {
				return err
			}

			// Beacon executable
			err = util.Fetch("https://storage.googleapis.com/marlin-artifacts/bin/beacon-"+runtime.GOOS+"-"+runtime.GOARCH, usr.HomeDir+"/.marlin/ctl/bin/beacon", usr.Username, true, false)
			if err != nil {
				return err
			}

			// Beacon config
			err = util.Fetch("https://storage.googleapis.com/marlin-artifacts/configs/beacon.conf", usr.HomeDir+"/.marlin/ctl/configs/beacon.conf", usr.Username, false, false)
			if err != nil {
				return err
			}

			err = util.TemplatePlace(
				usr.HomeDir+"/.marlin/ctl/configs/beacon.conf",
				"/etc/supervisor/conf.d/"+program+".conf",
				struct {
					Program, User, UserHome string
					DiscoveryAddr, HeartbeatAddr, BeaconAddr string
				} {
					program, usr.Username, usr.HomeDir, discovery_addr, heartbeat_addr, beacon_addr,
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
