package commands

import (
	"github.com/urfave/cli/v2"

	"marlin-cli/relay"
)

var Relay = cli.Command{
	Name:  "relay",
	Usage: "create, start or stop relay",
	Subcommands: []*cli.Command{
		relay.CreateCommand(),
		relay.DestroyCommand(),
		relay.StartCommand(),
		relay.StopCommand(),
		relay.RestartCommand(),
	},
}
