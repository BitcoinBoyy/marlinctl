package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"marlin-cli/relay"
)

var Relay = cli.Command{
	Name:  "relay",
	Usage: "create, start or stop relay",
	Subcommands: []*cli.Command{
		relay.CreateCommand(),
		relay.StartCommand(),
		{
			Name:  "stop",
			Usage: "stop the relay",
			Action: func(c *cli.Context) error {
				processName := "relay"
				if err := StopProcess(processName); err != nil {
					fmt.Println("error while stopping process: ", processName, err)
				}
				fmt.Println("stopped: ", processName)
				return nil
			},
		},
	},
}
