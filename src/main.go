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
	} else if os.Args[1] == "test" {
		cli.LevelTestStarter()
	} else if os.Args[1] == "nomadcoin" {
		//cli.NomadCoinStarter() //TODO : 노마드코인 수업 내용 붙이는 작업
	}
}
