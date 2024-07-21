package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mokhtarHamdoue/dmupdater/core"
)

func readStringFromBuffer(reader *bufio.Reader) (string, error) {
	value, err := reader.ReadString('\n')
	value = strings.Trim(value, "\n")
	return value, err
}

type CLIAPPCreator struct{}

func (cpc CLIAPPCreator) CreateAPP() *core.APP {
	reader := bufio.NewReader(os.Stdin)
	app := core.NewEmptyApplication()

	for {
		fmt.Print("Name of the application > ")
		appName, _ := readStringFromBuffer(reader)
		err := app.SetName(appName)
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

	for {
		fmt.Print("Version of the application > V")
		appVersion, _ := readStringFromBuffer(reader)
		err := app.SetVersion(appVersion)
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

	fmt.Print("Url of the repository >")
	url, _ := readStringFromBuffer(reader)
	app.SetUrl(url)

	fmt.Print("Is this version stable (Y/N), it's N if you skip > ")
	isLts, _ := readStringFromBuffer(reader)
	app.SetIsLTS(isLts == "Y")

	for {
		fmt.Print("Release date of this version (YYYY-MM-DD), it's today if you skip > ")
		releaseDate, _ := readStringFromBuffer(reader)
		if releaseDate != "" {
			err := app.SetReleaseDate(releaseDate)
			fmt.Println(err)
			continue
		} else {
			app.SetReleaseDate(time.Now().UTC().Format("2006-01-02"))
		}
		break
	}

	return app
}
