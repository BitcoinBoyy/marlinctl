package util

import(
	"os"
	"os/user"
)


func GetUser() (*user.User, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	if os.Geteuid() == 0 {
		// Root, try to retrieve SUDO_USER if exists
		if u := os.Getenv("SUDO_USER"); u != "" {
			usr, err = user.Lookup(u)
			if err != nil {
				return nil, err
			}
		}
	}

	return usr, nil
}
