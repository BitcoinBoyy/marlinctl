package util

import (
	"os"
	"os/exec"
	"net/http"
	"errors"
	"path/filepath"
	"io"
)

func Fetch(url, path, usr string, isExecutable bool, overwrite bool) error {
	// Check if already exists
	if !overwrite {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			return nil
		}
	}

	// Create dir
	_, err := exec.Command("sudo", "-u", usr, "mkdir", "-p", filepath.Dir(path)).Output()
	if err != nil {
		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("Fetch error")
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, resp.Body)
	f.Close()
	if err != nil {
		return err
	}

	// Perms
	_, err = exec.Command("chown", usr+":"+usr, path).Output()
	if err != nil {
		return err
	}

	if isExecutable {
		err = os.Chmod(path, 0755)
	} else {
		err = os.Chmod(path, 0644)
	}
	if err != nil {
		return err
	}

	return nil
}
