package logger

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func WriteToError(err error, exitTheApp bool) {

	if err == nil {
		return
	}

	file, _ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	text := fmt.Sprintf("[%s] INFO: %s\r", time.Now().Format(time.DateTime), err.Error())

	file.WriteString(text)

	if exitTheApp {
		os.Exit(1)
	}

}

func WriteAnyText(s string) {
	WriteToError(errors.New("NO ERROR - "+s), false)
}
