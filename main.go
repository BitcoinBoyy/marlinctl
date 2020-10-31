package main

import (
	"fmt"
	"log"
	"marlinctl/commands"
	"os"
)

func main() {

	if !commands.IsRoot() {
		fmt.Println("requires root permissions. Please run with sudo")
		return
	}
	if !commands.IsCommandAvailable("supervisorctl") {
		fmt.Println("supervisorctl not installed!!! Please install and try again")
		return
	}
	err := commands.App.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
