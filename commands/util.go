package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func IsProcessRunning(processName string) bool {
	out, err := exec.Command("sudo", "supervisorctl", "pid", processName).Output()

	if err != nil {
		fmt.Println("error is checking for process: ", processName, err)
		return false
	} else if string(out) == ("No such process "+processName) || string(out) == "0" {
		return false
	}
	return true
}

func StartProcess(processName string) error {
	out, err := exec.Command("sudo", "supervisorctl", "start", processName).Output()

	if err != nil {
		return err
	}
	fmt.Println("process started: ", processName, string(out))
	return nil
}

func UpdateRunningProcess(processName string) error {
	out, err := exec.Command("sudo", "supervisorctl", "update", processName).Output()

	if err != nil {
		return err
	}
	fmt.Println("process updated: ", processName, string(out))
	return nil
}

func StopProcess(processName string) error {
	out, err := exec.Command("sudo", "supervisorctl", "stop", processName).Output()

	if err != nil {
		return err
	}
	fmt.Println("process stopped: ", processName, string(out))
	return nil
}

func UpdateCommand(commandName string, newCommand string) error {
	configFilePath := "../supervisord.conf"
	newCommand = "command=" + newCommand
	_, err := exec.Command("sudo", "sed", "-i", "s+.*"+commandName+".*+"+newCommand+"+", configFilePath).Output()
	if err != nil {
		return err
	}
	return nil
}

func IsCommandAvailable(name string) bool {
	cmd := exec.Command("/bin/sh", "-c", "command -v "+name)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func IsRoot() bool {
	return os.Geteuid() == 0
}
