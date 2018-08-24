package loggenerator

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

func findOrCreateFile(fname string) (*os.File, error) {
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		_, err := os.Create(fname)
		if err != nil {
			return nil, err
		}
	}
	return os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
}

func TestFileWriting(t *testing.T) {
	randNumb := rand.Intn(100000)
	randomFileName := fmt.Sprintf("/tmp/random-%v-logs.txt", randNumb)
	file, err := os.Create(randomFileName)
	defer file.Close()
	if err != nil {
		t.Errorf("Error while creating file: %v", err)
	}

	lg := LogGenerator{
		Writer: file,
	}
	expected := 30

	for i := 0; i < expected; i++ {
		err = lg.Run()
		if err != nil {
			t.Errorf("Error while writing random log: %v", err)
		}
		time.Sleep(time.Millisecond)
	}
	f, err := os.Open(randomFileName)
	defer f.Close()
	if err != nil {
		t.Errorf("Error while opening file")
	}
	reader := bufio.NewScanner(f)
	linesCounter := 0
	for reader.Scan() {
		linesCounter++
	}
	if linesCounter != expected {
		t.Errorf("Error while file writing, expected lines: %v, result is %v", expected, linesCounter)
	}
}
