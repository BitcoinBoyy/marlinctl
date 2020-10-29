package commands

import (
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
		beacon.StopCommand(),
		beacon.RestartCommand(),
	},
}
