package core

type APP struct {
	name string
	version string
	url string
	release_note string
	isLTS  bool
	applications []APP
}

func (a *APP) setVersion(newVersion string){
	a.version = newVersion
}

func (a *APP) setURL(newURL string){
	a.url = newURL
}

func (a *APP) setReleaseNote(newRN string){
	a.release_note = newRN
}

func (a *APP) setIsLTS(isLts bool){
	a.isLTS = isLts
}

// add a sub application to the main application
func (a *APP) addAPP(newApp APP){
	a.applications = append(a.applications, newApp)
}

// delete an sub application from the main application
func (a *APP) deleteAto(appName string){
	var newAppList []APP = make([]APP, 0)
	for _, app  := range(a.applications) {
		if app.name != appName {
			newAppList = append(newAppList, app)
		}	
	}
	a.applications = newAppList 
}


