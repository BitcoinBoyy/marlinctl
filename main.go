package main

import (
	"log"
	"marlin-cli/commands"
	"os"
)

func main() {

	err := commands.App.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
