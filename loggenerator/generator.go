package loggenerator

import (
	"bufio"
	"fmt"
	"io"
	"time"

	"github.com/bxcodec/faker"
)

//LogGenerator struct
type LogGenerator struct {
	Writer io.Writer
}

//Run method for writing random log
func (lg *LogGenerator) Run() error {
	writer := bufio.NewWriter(lg.Writer)
	randString := lg.generateRandomLine()
	_, err := writer.WriteString(randString)
	if err == nil {
		writer.Flush()
	}
	return err
}

func (lg *LogGenerator) generateRandomLine() string {
	var randRecord string
	faker.FakeData(&randRecord)
	t := time.Now()
	return fmt.Sprintf("%v | %v\n", t.Format("Jan 2, 2006 at 3:04:05pm (UTC)"), randRecord)
}
