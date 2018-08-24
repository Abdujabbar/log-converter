package main

import (
	"fmt"
	"os"

	"github.com/Abdujabbor/log-converter/repository"
	"github.com/Abdujabbor/log-converter/watcher"
)

var tableTmpl = "<table border=1 cellpadding=10 cellspacing=5 style='width: 100%%'><thead><th>ID</th><th>TIME</th><th>MSG</th><th>FORMAT</th></thead><tbody>%s</tbody><tfoot>%s</tfoot></table>"
var rawTmpl = "<tr><td>%v</td><td>%v</td><td>%v</td><td>%v</td></tr>"
var tfootTpml = "<tr><td colspan=3>Total Rows</td><td>%v</td></tr>"

func main() {
	// f, err := findOrCreateFile("/tmp/random-logs.txt")

	// watcher.FileLinesFetcher(f)
	// fmt.Println("Please wait until servers will be ready")

	// provider := repository.Provider{
	// 	Server:   "localhost",
	// 	Database: "plogs",
	// }

	// err := provider.Connect()

	// if err != nil {
	// 	panic(err)
	// }

	// err = provider.Truncate()

	// if err != nil {
	// 	panic(err)
	// }

	// files, err := parseArgs(os.Args)

	// if err != nil {
	// 	panic(err)
	// }

	// go monitoringFiles(files, &provider)

	// router := handlers.InitRoutes(provider)

	// http.ListenAndServe(":8080", router)
}

func run(files []string, provider *repository.Provider) {
	// for _, v := range files {
	// 	go collectFileRecordsToDB(provider, v)
	// }
}

func parseArgs(args []string) ([]string, error) {
	if len(args) < 3 {
		return nil, fmt.Errorf("Required parameters files and format doesn't passed")
	}

	if _, ok := watcher.SupportedFormats[args[len(args)-1]]; !ok {
		return nil, fmt.Errorf("format parameter can receive only %v", watcher.SupportedFormats)
	}

	files := args[1 : len(args)-1]

	for _, v := range files {
		_, err := findOrCreateFile(v)
		if err != nil {
			panic(err)
		}
	}
	return files, nil
}

func findOrCreateFile(fname string) (*os.File, error) {
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		return os.Create(fname)
	}
	return os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
}
