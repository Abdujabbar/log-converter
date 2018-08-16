package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/Abdujabbor/log-converter/repository"
)

func TestProcess(t *testing.T) {
	dataAccessObject := repository.DAO{
		Server:   "localhost",
		Database: "testdb",
	}
	dataAccessObject.Connect()
	dataAccessObject.Truncate()

	randNumb := rand.Intn(100000)
	file := fmt.Sprintf("/tmp/logs-%s.txt", strconv.Itoa(randNumb))
	fmt.Printf("Random file name: %v\n", file)
	checkAndCreateFile(file)
	args := []string{
		"program",
		file,
		"1",
	}
	go process(&dataAccessObject, args)
	expected := 50
	for i := 0; i < expected; i++ {
		writeRandomLog(file)
		time.Sleep(time.Millisecond * 5)
	}

	records, err := dataAccessObject.FindAll(expected+1, 0)

	if err != nil {
		t.Error(err)
	}
	if len(records) != expected {
		t.Errorf("Error while storing on database, expected %v, received %v", expected, len(records))
	}
}
