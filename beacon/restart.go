package beacon

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func RestartCommand() *cli.Command {
	return &cli.Command{
		Name:  "restart",
		Usage: "restart the beacon",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			_, err := exec.Command("sudo", "supervisorctl", "restart", "beacon").Output()
			if err != nil {
				return err
			}

			output, _ := exec.Command("supervisorctl", "status").Output()
			fmt.Print(string(output))

			return nil
		},
	}
}
