package beacon

import (
	"os/exec"

	"github.com/urfave/cli/v2"
)


func StartCommand() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "start the beacon",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			_, err := exec.Command("sudo", "supervisorctl", "start", "beacon").Output()
			if err != nil {
				return err
			}

			return nil
		},
	}
}