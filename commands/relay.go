package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"marlin-cli/relay"
)

var Relay = cli.Command{
	Name:  "relay",
	Usage: "create, start or stop relay",
	Subcommands: []*cli.Command{
		relay.CreateCommand(),
		{
			Name:  "start",
			Usage: "start the relay",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "param1",
					Value:       "t",
					Usage:       "--param1 n",
					Destination: &param1,
				},
			},
			Action: func(c *cli.Context) error {
				processName := "relay"
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
			Usage: "stop the relay",
			Action: func(c *cli.Context) error {
				processName := "relay"
				if err := StopProcess(processName); err != nil {
					fmt.Println("error while stopping process: ", processName, err)
				}
				fmt.Println("stopped: ", processName)
				return nil
			},
		},
	},
}
