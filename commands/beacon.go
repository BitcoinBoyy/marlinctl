package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"marlin-cli/beacon"
)

var param1 string
var Beacon = cli.Command{

	Name:  "beacon",
	Usage: "create, start or stop beacon",
	Subcommands: []*cli.Command{
		beacon.CreateCommand(),
		beacon.StartCommand(),
		{
			Name:  "stop",
			Usage: "stop the beacon",
			Action: func(c *cli.Context) error {
				processName := "beacon"
				if err := StopProcess(processName); err != nil {
					fmt.Println("error while stopping process: ", processName, err)
				}
				fmt.Println("stopped: ", processName)
				return nil
			},
		},
	},
}
