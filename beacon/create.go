package beacon

import (
	"os"
	"net/http"
	"io"
	"errors"
	"path/filepath"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"

	"marlin-cli/util"
)


func CreateCommand() *cli.Command {
	var discovery_addr *string
	var heartbeat_addr *string
	var beacon_addr *string
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
				Destination: discovery_addr,
			},
			&cli.StringFlag{
				Name:        "heartbeat-addr",
				Usage:       "--heartbeat-addr <IP:PORT>",
				Destination: heartbeat_addr,
			},
			&cli.StringFlag{
				Name:        "beacon-addr",
				Usage:       "--beacon-addr <IP:PORT>",
				Destination: beacon_addr,
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
			err = fetch("https://storage.googleapis.com/marlin-artifacts/bin/beacon", usr.HomeDir+"/.marlin/ctl/bin/beacon", usr.Username, true)
			if err != nil {
				return err
			}

			// Beacon config
			err = fetch("https://storage.googleapis.com/marlin-artifacts/configs/beacon.conf", usr.HomeDir+"/.marlin/ctl/configs/beacon.conf", usr.Username, false)
			if err != nil {
				return err
			}

			err = util.TemplatePlace(
				usr.HomeDir+"/.marlin/ctl/configs/beacon.conf",
				"/etc/supervisor/conf.d/"+program+".conf",
				struct {
					Program, User, UserHome string
					DiscoveryAddr, HeartbeatAddr, BeaconAddr *string
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

func fetch(url, path, usr string, isExecutable bool) error {
	// Create dir
	_, err := exec.Command("sudo", "-u", usr, "mkdir", "-p", filepath.Dir(path)).Output()
	if err != nil {
		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("Fetch error")
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, resp.Body)
	f.Close()
	if err != nil {
		return err
	}

	// Perms
	_, err = exec.Command("chown", usr+":"+usr, path).Output()
	if err != nil {
		return err
	}

	if isExecutable {
		err = os.Chmod(path, 0755)
	} else {
		err = os.Chmod(path, 0644)
	}
	if err != nil {
		return err
	}

	return nil
}
