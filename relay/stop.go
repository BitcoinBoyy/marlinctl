package relay

import (
	"os/exec"

	"github.com/urfave/cli/v2"
)


func StopCommand() *cli.Command {
	return &cli.Command{
		Name:  "stop",
		Usage: "stop the relay",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			_, err := exec.Command("sudo", "supervisorctl", "stop", "relay").Output()
			if err != nil {
				return err
			}

			return nil
		},
	}
}
