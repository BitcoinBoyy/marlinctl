package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var App = &cli.App{
	Name:  "marlin-cli",
	Usage: "",
	Action: func(c *cli.Context) error {
		fmt.Println("type help to see usage...")
		return nil
	},
	Commands: []*cli.Command{
		&Beacon,
		&Relay,
	},
}
