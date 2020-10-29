package beacon

import (
	"text/template"
	"os"
	"os/user"
	"net/http"
	"io"
	"errors"
	"path/filepath"

	"github.com/urfave/cli/v2"
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
			usr, _ := user.Current()
			if os.Geteuid() == 0 {
				// Root, try to retrieve SUDO_USER if exists
				if u := os.Getenv("SUDO_USER"); u != "" {
					usr, _ = user.Lookup(u)
				}
			}

			err := fetch("https://storage.googleapis.com/marlin-artifacts/configs/beacon.conf", usr.HomeDir+"/.marlin/ctl/configs/beacon.conf")
			if err != nil {
				return err
			}
			err = fetch("https://storage.googleapis.com/marlin-artifacts/bin/beacon", usr.HomeDir+"/.marlin/ctl/bin/beacon")
			if err != nil {
				return err
			}

			t, err := template.ParseFiles(usr.HomeDir+"/.marlin/ctl/configs/beacon.conf")
			// t, err := template.ParseFiles("./configs/beacon.conf")
			if err != nil {
				return err
			}

			f, err := os.Create("/etc/supervisor/conf.d/beacon.conf")
			if err != nil {
				return err
			}

			err = t.Execute(f, struct {
				Program, User, UserHome string
				DiscoveryAddr, HeartbeatAddr, BeaconAddr *string
			} {
				"beacon", usr.Username, usr.HomeDir, discovery_addr, heartbeat_addr, beacon_addr,
			})
			f.Close()
			if err != nil {
				return err
			}

			out, err := exec.Command("sudo", "supervisorctl", "reread", "beacon").Output()
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
