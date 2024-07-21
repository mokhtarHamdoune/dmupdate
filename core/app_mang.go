package core

import "fmt"

// an interface you should inherit
type AppCreator interface {
	CreateAPP() *APP
}

type AppsManager struct {
	appCreator AppCreator
	apps       []APP
}

func NewManager(ac AppCreator) AppsManager {
	return AppsManager{
		appCreator: ac,
		apps:       make([]APP, 0),
	}
}

func (am *AppsManager) CreateNewApp() {
	newApp := am.appCreator.CreateAPP()
	am.apps = append(am.apps, *newApp)
	fmt.Println(newApp)
}
