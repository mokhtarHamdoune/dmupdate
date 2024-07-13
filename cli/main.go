package main

import (
	"fmt"
	"os"

	"github.com/mokhtarHamdoue/dmupdater/core"
)

type CliIO struct{}

func (c *CliIO) Write(p []byte) (int, error){
	return os.Stdout.Write(p)
}

func (c *CliIO) Read(p []byte) (int, error){
	return os.Stdin.Read(p)
}

func main(){
	
	appsManager := core.AppsManager{}

	appsManager.SetAppsManagerIO(&CliIO{})

	newApp, err := appsManager.CreateApplication()
	
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newApp)
}
