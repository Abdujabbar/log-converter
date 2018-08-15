package main

import (
	"fmt"
	"os"
)

func process(args []string) {
	_, files, err := parseInputArgs(args)
	if err != nil {
		panic(err)
	}

	for _, v := range files {
		go collectFileRecordsToDB(v)
	}
}

func parseInputArgs(args []string) ([]string, []string, error) {
	if len(args) < 3 {
		return nil, nil, fmt.Errorf("Required parameters files and format doesn't passed")
	}
	availableFormats := map[string]bool{}
	availableFormats[ftype] = true
	availableFormats[stype] = true
	if _, ok := availableFormats[args[len(args)-1]]; !ok {
		return nil, nil, fmt.Errorf("format parameter can receive only %v or %v", ftype, stype)
	}

	files := args[1 : len(args)-1]

	for _, v := range files {
		if _, err := os.Stat(v); os.IsNotExist(err) {
			os.Create(v)
		}
	}
	return args, files, nil
}
