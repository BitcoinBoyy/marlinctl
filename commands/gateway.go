package commands

import (
	"github.com/urfave/cli/v2"

	"marlinctl/dot_gateway"
)

var Gateway = cli.Command{
	Name:  "gateway",
	Usage: "create, start or stop gateways",
	Subcommands: []*cli.Command{
		&dot_gateway.DotGateway,
	},
}
