package filewatcher

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Abdujabbor/log-converter/repository"
	"github.com/bxcodec/faker"
	"github.com/fsnotify/fsnotify"
)

//SupportedFormats for parsing
var SupportedFormats = map[string]string{
	"human_readable": "Jan 2, 2006 at 3:04:05pm (UTC)",
	"rfc_1545":       "2006-02-01T15:04:05Z",
}

//FileSizes stores current size
type FileSizes map[string]int64

var fSizes = make(FileSizes)

//Start starting the process
func Start(recordChan chan *repository.Record, files []string) {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("error:", err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					line, err := getLastLine(event.Name)
					if err != nil {
						log.Println("error: ", err)
						continue
					}
					record := buildRecordFromString(line)
					recordChan <- record
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	for _, v := range files {
		err = watcher.Add(v)
		if err != nil {
			log.Println("error:", err)
		}
		info, err := os.Stat(v)
		if err != nil {
			log.Println("error:", err)
		}
		fSizes[v] = info.Size()
	}

	if err != nil {
		log.Println("error:", err)
	}
	<-done
}

func buildRecordFromString(line string) *repository.Record {
	r := repository.NewRecord()
	s := strings.Split(strings.Trim(line, "\n"), " | ")
	if len(s) < 2 {
		return nil
	}
	t, msg := strings.TrimSpace(s[0]), strings.TrimSpace(s[1])
	for k, v := range SupportedFormats {
		tm, err := time.Parse(v, t)
		if err != nil {
			continue
		}
		r.Format = k
		r.Time = tm.Unix()
		break
	}

	r.Msg = msg
	return r
}

func getLastLine(file string) (string, error) {
	info, err := os.Stat(file)
	if err != nil {
		return "", err
	}

	var oldSize int64
	if v, ok := fSizes[file]; ok {
		oldSize = v
	}

	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return "", err
	}
	fSizes[file] = info.Size()
	rbytes := make([]byte, info.Size()-oldSize)
	f.ReadAt(rbytes, oldSize)
	return string(rbytes), nil
}

func generateRandomLine() string {
	var randRecord string
	faker.FakeData(&randRecord)
	t := time.Now()
	return fmt.Sprintf("%v | %v\n", t.Format("Jan 2, 2006 at 3:04:05pm (UTC)"), randRecord)
}
