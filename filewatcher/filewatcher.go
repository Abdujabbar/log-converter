package filewatcher

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Abdujabbor/log-converter/repository"
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
func Start(provider *repository.Provider, files []string) {
	done := make(chan bool)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("error:", err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					line, err := readLastLine(event.Name)
					if err != nil {
						log.Println("error: ", err)
						continue
					}
					record := buildRecordFromLine(line)
					err = provider.Insert(record)
					if err != nil {
						log.Println("error: ", err)
					} else {
						records, err := provider.FindAll(-1, 0)
						if err != nil {
							log.Println("error: ", err)
						} else {
							log.Println("Stored records count:", len(records))
						}
					}
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

func buildRecordFromLine(line string) *repository.Record {
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

func readLastLine(fname string) (string, error) {
	info, err := os.Stat(fname)
	if err != nil {
		return "", err
	}
	var oldSize int64
	if v, ok := fSizes[fname]; ok {
		oldSize = v
	}
	f, err := os.Open(fname)

	if err != nil {
		return "", err
	}

	fSizes[fname] = info.Size()

	f.Seek(oldSize, int(fSizes[fname]-oldSize))

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	return scanner.Text(), nil
}
