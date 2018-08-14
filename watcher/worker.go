package watcher

import (
	"fmt"
	"os"
	"time"
)

//FRange struct
type FRange struct {
	Start int64
	End   int64
}

//Process method
func Process(frangeChan chan FRange, fname string) {
	var fsize int64
	fsize = -1
	for {
		finfo, err := os.Stat(fname)
		if err != nil {
			fmt.Println(err)
		}
		if fsize == -1 {
			fsize = finfo.Size()
		} else {
			if finfo.Size() != fsize {
				fr := FRange{
					Start: fsize,
					End:   finfo.Size(),
				}
				frangeChan <- fr
				fsize = finfo.Size()
			}
		}
		time.Sleep(time.Millisecond * 100)
	}

}
