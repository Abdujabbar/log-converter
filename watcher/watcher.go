package watcher

import (
	"bufio"
	"fmt"
	"io"
)

//SupportFormats var
var SupportedFormats = map[string]string{
	"human_readable": "Jan 2, 2006 at 3:04:05pm (UTC)",
	"rfc_1545":       "2006-02-01T15:04:05Z",
}

func FileLinesFetcher(reader io.Reader) {
	bufferedReader := bufio.NewScanner(reader)

	for bufferedReader.Scan() {
		fmt.Printf("Scanned text: %v\n", bufferedReader.Text())
	}
}

//FileWriteNotifiy method
// func FileWriteNotifiy(recordChan chan *repository.Record, fname string) {
// 	var fsize int64
// 	fsize = -1
// 	file, _ := os.Open(fname)
// 	defer file.Close()

// 	for {
// 		finfo, err := os.Stat(fname)
// 		if err != nil {
// 			fmt.Println(err)
// 			break
// 		}
// 		if fsize == -1 {
// 			fsize = finfo.Size()
// 		} else {
// 			if finfo.Size() != fsize {
// 				rbytes := make([]byte, finfo.Size()-fsize)
// 				file.ReadAt(rbytes, fsize)
// 				record := parseRaw(string(rbytes))
// 				if record != nil {
// 					record.FileName = fname
// 					recordChan <- record
// 				}
// 				fsize = finfo.Size()
// 			}
// 		}
// 		time.Sleep(time.Nanosecond)
// 	}
// }

// func parseRaw(raw string) *repository.Record {
// 	r := repository.NewRecord()
// 	// s := strings.Split(strings.Trim(raw, "\n"), " | ")
// 	// t, msg := strings.TrimSpace(s[0]), strings.TrimSpace(s[1])
// 	// tm, err := time.Parse(firstTimeFormat, t)
// 	// r.Format = "first_format"
// 	// if err != nil {
// 	// 	tm, err = time.Parse(secondTimeFormat, t)
// 	// 	if err != nil {
// 	// 		return nil
// 	// 	}
// 	// 	r.Format = "second_format"
// 	// }
// 	// r.Time = tm.Unix()
// 	// r.Msg = msg
// 	return r
// }
