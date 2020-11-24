package relay

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func StopCommand() *cli.Command {
	var chain string

	return &cli.Command{
		Name:  "stop",
		Usage: "stop the relay",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "chain",
				Usage:       "--chain \"<CHAIN>\"",
				Destination: &chain,
				Required:    true,
			},
		},
		Action: func(c *cli.Context) error {
			program := chain + "_relay"

			_, err := exec.Command("sudo", "supervisorctl", "stop", program).Output()
			if err != nil {
				return err
			}

			output, _ := exec.Command("supervisorctl", "status").Output()
			fmt.Print(string(output))

			return nil
		},
	}
}
