package commands

import (
	"github.com/urfave/cli/v2"

	"marlinctl/gateways/dotGateway"
	"marlinctl/gateways/irisGateway"
	"marlinctl/gateways/nearGateway"
)

var Gateway = cli.Command{
	Name:  "gateway",
	Usage: "create, start or stop gateways",
	Subcommands: []*cli.Command{
		&dotGateway.DotGateway,
		&irisGateway.IrisGateway,
		&nearGateway.NearGateway,
	},
}
