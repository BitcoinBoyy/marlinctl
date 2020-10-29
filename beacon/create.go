package beacon

import (
	"os"
	"net/http"
	"io"
	"errors"
	"path/filepath"
	"os/exec"

	"github.com/urfave/cli/v2"

	"marlin-cli/util"
)


func CreateCommand() *cli.Command {
	var discovery_addr *string
	var heartbeat_addr *string
	var beacon_addr *string

	return &cli.Command{
		Name:  "create",
		Usage: "create a new beacon",
		Flags: []cli.Flag{
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
			// User details
			usr, err := util.GetUser()
			if err != nil {
				return err
			}

			// Beacon executable
			err = fetch("https://storage.googleapis.com/marlin-artifacts/bin/beacon", usr.HomeDir+"/.marlin/ctl/bin/beacon")
			if err != nil {
				return err
			}

			// Beacon config
			err = fetch("https://storage.googleapis.com/marlin-artifacts/configs/beacon.conf", usr.HomeDir+"/.marlin/ctl/configs/beacon.conf")
			if err != nil {
				return err
			}

			err = util.TemplatePlace(
				usr.HomeDir+"/.marlin/ctl/configs/beacon.conf",
				"/etc/supervisor/conf.d/beacon.conf",
				struct {
					Program, User, UserHome string
					DiscoveryAddr, HeartbeatAddr, BeaconAddr *string
				} {
					"beacon", usr.Username, usr.HomeDir, discovery_addr, heartbeat_addr, beacon_addr,
				},
			)
			if err != nil {
				return err
			}

			_, err = exec.Command("sudo", "supervisorctl", "reread", "beacon").Output()
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func fetch(url, path string) error {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
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
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
