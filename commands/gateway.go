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
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "param1",
					Value:       "t",
					Usage:       "--param1 n",
					Destination: &param1,
				},
			},
			Action: func(c *cli.Context) error {
				processName := "gateway"
				sampleCommand := "/bin/cat"
				UpdateCommand(sampleCommand, sampleCommand+" -"+param1)
				if IsProcessRunning(processName) {
					if err := UpdateRunningProcess(processName); err != nil {
						fmt.Println("error while starting process: ", processName, err)
					}
				} else {
					if err := StartProcess(processName); err != nil {
						fmt.Println("error while starting process: ", processName, err)
					}
				}
				fmt.Println("started: ", processName)
				return nil
			},
		},
		{
			Name:  "stop",
			Usage: "stop the gateway",
			Action: func(c *cli.Context) error {
				processName := "gateway"
				if err := StopProcess(processName); err != nil {
					fmt.Println("error while stopping process: ", processName, err)
				}
				fmt.Println("stopped: ", processName)
				return nil
			},
		},
	},
}
