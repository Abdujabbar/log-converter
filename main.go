package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Abdujabbor/log-converter/filewatcher"
	"github.com/Abdujabbor/log-converter/repository"
	// "github.com/Abdujabbor/log-converter/watcher"
)

// var tableTmpl = "<table border=1 cellpadding=10 cellspacing=5 style='width: 100%%'><thead><th>ID</th><th>TIME</th><th>MSG</th><th>FORMAT</th></thead><tbody>%s</tbody><tfoot>%s</tfoot></table>"
// var rawTmpl = "<tr><td>%v</td><td>%v</td><td>%v</td><td>%v</td></tr>"
// var tfootTpml = "<tr><td colspan=3>Total Rows</td><td>%v</td></tr>"
//FileSize stores current size

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

	if err != nil {
		log.Println("error", err)
	}

	filewatcher.Start(&provider, files)
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
