package tmGateway

import (
	"github.com/urfave/cli/v2"
)

var TmGateway = cli.Command{
	Name:  "tendermint",
	Usage: "create, start or stop tendermint gateways",
	Subcommands: []*cli.Command{
		CreateCommand(),
		DestroyCommand(),
	},
}
