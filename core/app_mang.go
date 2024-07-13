package core

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

type AppsManangerIO interface {
	 io.Reader
	 io.Writer
}

type AppsManager struct {
	amIO AppsManangerIO
}

func (am * AppsManager) SetAppsManagerIO (amio AppsManangerIO){
		am.amIO = amio	
}


func (am * AppsManager) CreateApplication() (*APP, error) {

	if am.amIO == nil {
		return nil, errors.New("AppsManangerIO is not provided!")
	}

	_, err := am.amIO.Write([]byte("DM Update V0.1.0\n"))
	_, err = am.amIO.Write([]byte("Name of application > "))
	appName, err :=  am.readFromBuffer()
	_, err = am.amIO.Write([]byte("Version of application > V"))
	appVersion, err :=  am.readFromBuffer()


	if err != nil {
		fmt.Println("we sucesseed I think")
	}

	return &APP{
		name: appName, 
		version: "V"+appVersion, 
		isLTS: true, 
		url: "",
		release_note: "",
	}, nil
}

func (am * AppsManager) readFromBuffer() (string, error){
	// I used bufio because it give me the flexiblity to read string 
	// and to exit when the \n is found which is enter 
	// using am.amIO.Read to read the buffer does not give EOF to exit
	reader := bufio.NewReader(am.amIO)
	line , err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	line = strings.TrimSpace(line)
	return line, nil
}


