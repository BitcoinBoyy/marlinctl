package util

import (
	"encoding/json"
	"errors"
	"io"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func Fetch(url, path, usr string, isExecutable bool, overwrite bool) error {
	fmt.Println("Fetching", url, path, usr)
	// Check if already exists
	if !overwrite {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			return nil
		}
	}

	// Create dir
	_, err := exec.Command("mkdir", "-p", filepath.Dir(path)).Output()
	if err != nil {
		return err
	}
	fmt.Println("DBG")
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("DBG")

	if resp.StatusCode != 200 {
		return errors.New("Fetch error")
	}
	fmt.Println("DBG")

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	fmt.Println("DBG")

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

func FetchLatestVersion(configName string) (string, error) {

	resp, err := http.Get("https://storage.googleapis.com/marlin-artifacts/configs/versions.json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New("Fetch error")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	val, found := result[configName]
	if !found {
		return "", errors.New("Invalid config name")
	}
	return val.(string), nil
}
