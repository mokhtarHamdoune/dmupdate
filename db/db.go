package db

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/mokhtarHamdoue/dmupdater/core"
)

const DB_PATH string = "./db.json"
const ROOT_PROJECT_NAME = "dmupdater"

type JsonDB struct {
	db *os.File
}

func NewDB() (core.AppsDB, error) {
	file, err := createFileIFNotExist()

	if err != nil {
		return nil, err
	}

	return JsonDB{
		db: file,
	}, nil
}

// saveApp implements core.AppsDB.
func (j JsonDB) SaveApp(app core.APP) {
	panic("unimplemented")
}

func createFileIFNotExist() (*os.File, error) {
	db_path := getDBPath()
	file, err := os.Open(db_path)

	if err == nil {
		return file, nil
	}

	file, err = os.Create(db_path)
	if err != nil {
		return nil, errors.New("we could not create the data base")
	}

	return file, nil

}

func getDBPath() string {
	pwd, _ := os.Getwd()
	rootDirIndex := strings.LastIndex(pwd, ROOT_PROJECT_NAME)
	rootProjectPath := pwd[0 : rootDirIndex+len(ROOT_PROJECT_NAME)]
	db_path := filepath.Join(rootProjectPath, "db", DB_PATH)
	return db_path
}
