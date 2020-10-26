package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Gateway = cli.Command{

	Name:  "gateway",
	Usage: "create, start or stop gateway",
	Subcommands: []*cli.Command{
		{
			Name:  "create",
			Usage: "create a new gateway",
			Action: func(c *cli.Context) error {
				fmt.Println("created: ", c.Args().First())
				return nil
			},
		},
		{
			Name:  "start",
			Usage: "start the gateway",
			Action: func(c *cli.Context) error {
				fmt.Println("started: ", c.Args().First())
				return nil
			},
		},
		{
			Name:  "stop",
			Usage: "stop the gateway",
			Action: func(c *cli.Context) error {
				fmt.Println("stopped: ", c.Args().First())
				return nil
			},
		},
	},
}
