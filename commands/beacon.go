package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"marlin-cli/beacon"
)

var param1 string
var Beacon = cli.Command{

	Name:  "beacon",
	Usage: "create, start or stop beacon",
	Subcommands: []*cli.Command{
		beacon.CreateCommand(),
		{
			Name:  "start",
			Usage: "start the beacon",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "param1",
					Value:       "t",
					Usage:       "--param1 n",
					Destination: &param1,
				},
			},
			Action: func(c *cli.Context) error {
				processName := "beacon"
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
			Usage: "stop the beacon",
			Action: func(c *cli.Context) error {
				processName := "beacon"
				if err := StopProcess(processName); err != nil {
					fmt.Println("error while stopping process: ", processName, err)
				}
				fmt.Println("stopped: ", processName)
				return nil
			},
		},
	},
}
