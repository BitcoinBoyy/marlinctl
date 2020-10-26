package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Relay = cli.Command{

	Name:  "relay",
	Usage: "create, start or stop relay",
	Subcommands: []*cli.Command{
		{
			Name:  "create",
			Usage: "create a new relay",
			Action: func(c *cli.Context) error {
				fmt.Println("created: ", c.Args().First())
				return nil
			},
		},
		{
			Name:  "start",
			Usage: "start the relay",
			Action: func(c *cli.Context) error {
				fmt.Println("started: ", c.Args().First())
				return nil
			},
		},
		{
			Name:  "stop",
			Usage: "stop the relay",
			Action: func(c *cli.Context) error {
				fmt.Println("stopped: ", c.Args().First())
				return nil
			},
		},
	},
}
