package relay

import (
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)


func DestroyCommand() *cli.Command {
	return &cli.Command{
		Name:  "destroy",
		Usage: "destroy the relay",
		Flags: []cli.Flag{},
		Action: func(c *cli.Context) error {
			out, _ := exec.Command("sudo", "supervisorctl", "status", "relay").Output()
			if strings.Contains(string(out), "no such process") {
				// Already destroyed
				return nil
			}

			_, err := exec.Command("sudo", "supervisorctl", "stop", "relay").Output()
			if err != nil {
				return err
			}

			_, err = exec.Command("sudo", "supervisorctl", "remove", "relay").Output()
			if err != nil {
				return err
			}

			_, err = exec.Command("sudo", "rm", "/etc/supervisor/conf.d/relay.conf").Output()
			if err != nil {
				return err
			}

			_, err = exec.Command("supervisorctl", "reread").Output()
			if err != nil {
				return err
			}

			return nil
		},
	}
}
