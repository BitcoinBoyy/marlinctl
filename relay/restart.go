package relay

import (
	"os/exec"

	"github.com/urfave/cli/v2"
)


func RestartCommand() *cli.Command {
	return &cli.Command{
		Name:  "restart",
		Usage: "restart the relay",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			_, err := exec.Command("sudo", "supervisorctl", "restart", "relay").Output()
			if err != nil {
				return err
			}

			return nil
		},
	}
}
