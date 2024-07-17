package core

import (
	"errors"
	"regexp"
	"time"
)

type APP struct {
	name         string
	version      string
	url          string
	release_note string
	description  string
	isLTS        bool
	releaseDate  string
	applications []APP
}

func NewApplication(name string, version string) (*APP, error) {
	newApp := &APP{}

	err := newApp.SetName(name)
	if err != nil {
		return nil, err
	}

	err = newApp.SetVersion(version)
	if err != nil {
		return nil, err
	}

	newApp.SetUrl("")
	newApp.SetReleaseDate(time.Now().Format("YYYY-MM-DD"))
	newApp.SetReleaseNote("")
	newApp.SetDescription("")
	newApp.SetIsLTS(false)

	return newApp, nil
}
func NewEmptyApplication() *APP {
	return &APP{}
}

// name of the application (required)
// it return a panic if the app name is empty
func (a *APP) SetName(name string) error {
	if name == "" {
		return errors.New("application name should not be empty")
	}
	a.name = name
	return nil

}

// version of the application (required)
// it return a panic if the verion is empty or is not semantic version
func (a *APP) SetVersion(version string) error {

	if version == "" {
		return errors.New("application version is required")
	}

	isMatch, err := regexp.MatchString("^\\d+\\.\\d+\\.\\d+$", version)
	if err == nil && !isMatch {
		return errors.New("version must follow semantic version")
	}

	a.version = version

	return nil
}

// repo url current version of the application (optional)
func (a *APP) SetUrl(url string) {
	a.url = url
}

// release not of the current version of the application (optional)
func (a *APP) SetReleaseNote(newRN string) {
	a.release_note = newRN
}

// whether  the current version is stable of not (optional)
func (a *APP) SetIsLTS(isLts bool) {
	a.isLTS = isLts
}

// description of the application (optional)
func (a *APP) SetDescription(description string) {
	a.description = description
}

// release date of  of current version of the application (optional)
func (a *APP) SetReleaseDate(releaseDate string) error {
	rd, err := time.Parse("YYYY-MM-DD", releaseDate)
	if err != nil {
		return err
	}
	a.releaseDate = rd.String()
	return nil
}

// add a sub application to the main application
func (a *APP) AddAPP(newApp APP) {
	a.applications = append(a.applications, newApp)
}

// delete a sub application from the main application
func (a *APP) DeleteApp(appName string) {
	var newAppList []APP = make([]APP, 0)
	for _, app := range a.applications {
		if app.name != appName {
			newAppList = append(newAppList, app)
		}
	}
	a.applications = newAppList
}
