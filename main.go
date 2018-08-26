package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Abdujabbor/log-converter/filewatcher"
	"github.com/Abdujabbor/log-converter/handlers"
	"github.com/Abdujabbor/log-converter/repository"
)

func main() {

	files, err := parseArgs(os.Args)

	if err != nil {
		panic(err)
	}

	provider := repository.Provider{
		Server:   "localhost",
		Database: "test",
	}

	err = provider.Connect()

	if err != nil {
		panic(err)
	}

	err = provider.Truncate()

	if err != nil {
		panic(err)
	}

	router := handlers.InitRouter(provider)

	go http.ListenAndServe(":8080", router)

	recordChan := make(chan *repository.Record)

	go filewatcher.Start(recordChan, files)

	for {
		select {
		case record, ok := <-recordChan:
			if ok {
				err = provider.Insert(record)
				if err != nil {
					log.Println("error: ", err)
				}
			}
			break
		}
	}
}

func parseArgs(args []string) ([]string, error) {
	if len(args) < 3 {
		return nil, fmt.Errorf("Required parameters files and format doesn't passed")
	}

	if _, ok := filewatcher.SupportedFormats[args[len(args)-1]]; !ok {
		return nil, fmt.Errorf("format parameter can receive only %v", filewatcher.SupportedFormats)
	}

	files := args[1 : len(args)-1]

	for _, v := range files {
		err := findOrCreateFile(v)
		if err != nil {
			panic(err)
		}
	}
	return files, nil
}

func findOrCreateFile(fname string) error {
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		_, err = os.Create(fname)
		return err
	}
	return nil
}
