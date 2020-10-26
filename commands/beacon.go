package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Beacon = cli.Command{

	Name:  "beacon",
	Usage: "create, start or stop beacon",
	Subcommands: []*cli.Command{
		{
			Name:  "create",
			Usage: "create a new beacon",
			Action: func(c *cli.Context) error {
				fmt.Println("created: ", c.Args().First())
				return nil
			},
		},
		{
			Name:  "start",
			Usage: "start the beacon",
			Action: func(c *cli.Context) error {
				fmt.Println("started: ", c.Args().First())
				return nil
			},
		},
		{
			Name:  "stop",
			Usage: "stop the beacon",
			Action: func(c *cli.Context) error {
				fmt.Println("stopped: ", c.Args().First())
				return nil
			},
		},
	},
}
