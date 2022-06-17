package main

import (
	"os"
	"runner/cli"
)

func main() {
	if len(os.Args) < 2 {
		cli.Usage()
	} else if os.Args[1] == "edu" {
		cli.EduStarter()
	}
}
