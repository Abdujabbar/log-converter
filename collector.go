package main

import (
	"context"

	// "github.com/Abdujabbor/log-converter/notifier"
	"github.com/Abdujabbor/log-converter/repository"
)

func collectFileRecordsToDB(ctx context.Context, provider *repository.Provider, fname string) {
	// recordChan := make(chan *repository.Record)
	// go func(rChan chan *repository.Record) {
	// 	notifier.FileWriteNotifiy(rChan, fname)
	// }(recordChan)
	// for {
	// 	select {
	// 	case v, ok := <-recordChan:
	// 		if ok {
	// 			err := provider.Insert(*v)
	// 			if err != nil {
	// 				fmt.Printf("Error while storing record: %v, with string %v", v, err.Error())
	// 			}
	// 		}
	// 		break
	// 	}
	// }
}
