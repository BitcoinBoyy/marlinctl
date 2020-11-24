package irisGateway

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/urfave/cli/v2"

	"marlinctl/util"
)

func CreateCommand() *cli.Command {
	var bootstrapAddr, listenPortPeer, peerIP, peerPort, rpcPort string
	var version string

	return &cli.Command{
		Name:  "create",
		Usage: "create a new gateway",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "bootstrapaddr",
				Usage:       "--bootstrapaddr \"<IP1:PORT1>\"",
				Destination: &bootstrapAddr,
				Value:       "127.0.0.1:8002",
			},
			&cli.StringFlag{
				Name:        "listenportpeer",
				Usage:       "--listenportpeer \"PORT\"",
				Destination: &listenPortPeer,
				Value:       "59001",
			},
			&cli.StringFlag{
				Name:        "peerip",
				Usage:       "--peerip \"IP\"",
				Destination: &listenPortPeer,
				Value:       "127.0.0.1",
			},
			&cli.StringFlag{
				Name:        "peerport",
				Usage:       "--peerport \"PORT\"",
				Destination: &peerPort,
				Value:       "26656",
			},
			&cli.StringFlag{
				Name:        "rpcport",
				Usage:       "--rpcport \"PORT\"",
				Destination: &rpcPort,
				Value:       "26657",
			},
			&cli.StringFlag{
				Name:        "version",
				Usage:       "--version <NUMBER>",
				Value:       "latest",
				Destination: &version,
			},
		},
		Action: func(c *cli.Context) error {
			chain := "iris"
			program := chain + "_gateway"
			bridge_program := chain + "_bridge"
			keyfile := chain + "_keyfile"

			out, _ := exec.Command("sudo", "supervisorctl", "status", program).Output()
			if !strings.Contains(string(out), "no such process") {
				// Already exists
				return errors.New("Already exists")
			}

			// User details
			usr, err := util.GetUser()
			if err != nil {
				return err
			}

			// Version
			if version == "latest" {
				fmt.Println(program, "fetching latest binaries...")
				latestVersion, err := util.FetchLatestVersion(program)
				if err != nil {
					return err
				}
				version = latestVersion
				fmt.Println(program, "latest binary version: ", latestVersion)
			}

			// gateway executable
			err = util.Fetch("https://storage.googleapis.com/marlin-artifacts/bin/"+program+"-"+version+"-"+runtime.GOOS+"-"+runtime.GOARCH, usr.HomeDir+"/.marlin/ctl/bin/"+program+"-"+version, usr.Username, true, false)
			if err != nil {
				fmt.Println("Caught here")
				return err
			}

			// gateway config
			err = util.Fetch("https://storage.googleapis.com/marlin-artifacts/configs/"+program+"-"+version+".conf", usr.HomeDir+"/.marlin/ctl/configs/"+program+"-"+version+".conf", usr.Username, false, false)
			if err != nil {
				return err
			}

			// gateway keyfile
			err = util.Fetch("https://storage.googleapis.com/marlin-artifacts/configs/"+keyfile+"-"+version+".json", usr.HomeDir+"/.marlin/ctl/configs/"+keyfile+"-"+version+".json", usr.Username, false, false)
			if err != nil {
				return err
			}

			err = util.TemplatePlace(
				usr.HomeDir+"/.marlin/ctl/configs/"+program+"-"+version+".conf",
				"/etc/supervisor/conf.d/"+program+".conf",
				struct {
					Program, User, UserHome string
					GatewayVersion, KeyfileVersion string
					Listenportpeer, Peerip string
					Peerport, Rpcport      string
				}{
					program, usr.Username, usr.HomeDir,
					version, version,
					listenPortPeer, peerIP,
					peerPort, rpcPort,
				},
			)
			if err != nil {
				return err
			}

			// bridge executable
			err = util.Fetch("https://storage.googleapis.com/marlin-artifacts/bin/"+bridge_program+"-"+version+"-"+runtime.GOOS+"-"+runtime.GOARCH, usr.HomeDir+"/.marlin/ctl/bin/"+bridge_program+"-"+version, usr.Username, true, false)
			if err != nil {
				return err
			}

			// bridge config
			err = util.Fetch("https://storage.googleapis.com/marlin-artifacts/configs/"+bridge_program+"-"+version+".conf", usr.HomeDir+"/.marlin/ctl/configs/"+bridge_program+"-"+version+".conf", usr.Username, false, false)
			if err != nil {
				return err
			}

			err = util.TemplatePlace(
				usr.HomeDir+"/.marlin/ctl/configs/"+bridge_program+"-"+version+".conf",
				"/etc/supervisor/conf.d/"+bridge_program+".conf",
				struct {
					Program, User, UserHome string
					BootstrapAddr           string
					Version                 string
				}{
					bridge_program, usr.Username, usr.HomeDir,
					bootstrapAddr,
					version,
				},
			)
			if err != nil {
				return err
			}

			_, err = exec.Command("supervisorctl", "reread").Output()
			if err != nil {
				return err
			}

			_, err = exec.Command("supervisorctl", "add", program).Output()
			if err != nil {
				return err
			}

			_, err = exec.Command("supervisorctl", "add", bridge_program).Output()
			if err != nil {
				return err
			}

			output, _ := exec.Command("supervisorctl", "status").Output()
			fmt.Print(string(output))

			return nil
		},
	}
}
