package main

import (
	"log"

	"github.com/mokhtarHamdoue/dmupdater/core"
	"github.com/mokhtarHamdoune/dmupdate/db"
)

func main() {
	cliAppCreate := CLIAPPCreator{}
	jsonDB, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	applicationManger := core.NewManager(cliAppCreate, jsonDB)
	applicationManger.CreateNewApp()
}
