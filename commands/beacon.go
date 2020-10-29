package commands

import (
	"github.com/urfave/cli/v2"

	"marlin-cli/beacon"
)

var Beacon = cli.Command{
	Name:  "beacon",
	Usage: "create, start or stop beacon",
	Subcommands: []*cli.Command{
		beacon.CreateCommand(),
		beacon.DestroyCommand(),
		beacon.StartCommand(),
		beacon.StopCommand(),
		beacon.RestartCommand(),
	},
}
