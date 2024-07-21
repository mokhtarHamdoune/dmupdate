package main

import (
	"github.com/mokhtarHamdoue/dmupdater/core"
)

func main() {
	cliAppCreate := CLIAPPCreator{}
	applicationManger := core.NewManager(cliAppCreate)
	applicationManger.CreateNewApp()
}
