package filewatcher

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"testing"
)

func TestGetListLine(t *testing.T) {
	randNumb := rand.Int()
	randomFile := fmt.Sprintf("/tmp/random-%v.txt", randNumb)

	file, err := os.Create(randomFile)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	for i := 0; i < 5; i++ {
		expected := generateRandomLine()
		_, err = file.WriteString(expected)
		if err != nil {
			t.Error(err)
		}

		res, err := getLastLine(randomFile)
		if err != nil {
			t.Error(err)
		}
		res = strings.Trim(res, "\n")
		expected = strings.Trim(expected, "\n")
		if res != expected {
			t.Errorf("Expected %v, Received: %v", expected, res)
		}
	}
}
