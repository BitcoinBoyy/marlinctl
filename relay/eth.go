package relay

import (
	"os/exec"
	"runtime"
	"errors"
	"strings"

	"marlin-cli/util"
)

type EthAbci struct {}

func (abci *EthAbci) Create(datadir string) error {
	// User details
	usr, err := util.GetUser()
	if err != nil {
		return err
	}

	program := "geth"

	// geth executable
	err = util.Fetch("https://storage.googleapis.com/marlin-artifacts/bin/"+program+"-"+runtime.GOOS+"-"+runtime.GOARCH, usr.HomeDir+"/.marlin/ctl/bin/"+program, usr.Username, true, false)
	if err != nil {
		return err
	}

	// geth config
	err = util.Fetch("https://storage.googleapis.com/marlin-artifacts/configs/"+program+".conf", usr.HomeDir+"/.marlin/ctl/configs/"+program+".conf", usr.Username, false, false)
	if err != nil {
		return err
	}

	err = util.TemplatePlace(
		usr.HomeDir+"/.marlin/ctl/configs/"+program+".conf",
		"/etc/supervisor/conf.d/"+program+".conf",
		struct {
			Program, User, UserHome, DataDir string
		} {
			program, usr.Username, usr.HomeDir, datadir,
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

	return nil
}

func (abci *EthAbci) Destroy() error {
	program := "geth"

	out, _ := exec.Command("sudo", "supervisorctl", "status", program).Output()
	if strings.Contains(string(out), "no such process") {
		return errors.New("Not found")
	}

	_, err := exec.Command("sudo", "supervisorctl", "stop", program).Output()
	if err != nil {
		return err
	}

	_, err = exec.Command("sudo", "supervisorctl", "remove", program).Output()
	if err != nil {
		return err
	}

	_, err = exec.Command("sudo", "rm", "/etc/supervisor/conf.d/"+program+".conf").Output()
	if err != nil {
		return err
	}

	_, err = exec.Command("supervisorctl", "reread").Output()
	if err != nil {
		return err
	}

	return nil
}
