package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bxcodec/faker"
)

type record struct {
	Raw string
}

func writeRandomLog(fname string) error {
	var randRecord record
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		return err
	}
	faker.FakeData(&randRecord)
	loc, _ := time.LoadLocation("Asia/Tashkent")
	t := time.Now().In(loc)
	_, err = f.WriteString(fmt.Sprintf("%v | %v\n", t.Format("Jan 2, 2006 at 3:04:05pm (UTC)"), randRecord.Raw))
	if err != nil {
		return err
	}
	return nil
}
