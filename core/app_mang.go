package core

import "fmt"

// TODO: list the created app
// TODO: the list should be in presistand db
// TODO: create presistance storage it does not matter what kind of db you use, we should follow the rule that db is details
// TODO: when the manager start up the application should be retreived from the db or at least when it was asked

// an interface you should inherit
type AppCreator interface {
	CreateAPP() *APP
}

type AppsManager struct {
	appCreator AppCreator
	apps       []APP
	db         AppsDB
}

type AppsDB interface {
	SaveApp(app APP)
}

func NewManager(ac AppCreator, db AppsDB) AppsManager {
	return AppsManager{
		appCreator: ac,
		apps:       make([]APP, 0),
		db:         db,
	}
}

func (am *AppsManager) CreateNewApp() {
	newApp := am.appCreator.CreateAPP()
	am.db.SaveApp(*newApp)
	am.apps = append(am.apps, *newApp)
}

func (am *AppsManager) ListApplications() {
	for _, app := range am.apps {
		fmt.Println(app)
	}
}
