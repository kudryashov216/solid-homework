package logger

import (
	"errors"
	"os"
	"time"
)

func WriteToError(err error, exitTheApp bool) {

	if err == nil {
		return
	}

	file, _ := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	file.Close()

	file.Write([]byte("[" + time.Now().Format(time.DateTime) + "] INFO: " + err.Error()))

	if exitTheApp {
		os.Exit(1)
	}

}

func WriteAnyText(s string) {
	WriteToError(errors.New("NO ERROR - "+s), false)
}
