package beacon

import (
	"os/exec"

	"github.com/urfave/cli/v2"
)


func DestroyCommand() *cli.Command {
	return &cli.Command{
		Name:  "destroy",
		Usage: "destroy the beacon",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			_, err := exec.Command("sudo", "supervisorctl", "stop", "beacon").Output()
			if err != nil {
				return err
			}

			_, err = exec.Command("sudo", "supervisorctl", "remove", "beacon").Output()
			if err != nil {
				return err
			}

			_, err = exec.Command("sudo", "rm", "/etc/supervisor/conf.d/beacon.conf").Output()
			if err != nil {
				return err
			}

			_, err = exec.Command("supervisorctl", "reread").Output()
			if err != nil {
				return err
			}

			return nil
		},
	}
}
