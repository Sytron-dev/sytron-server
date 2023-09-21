package logger

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/withmandala/go-log"
)

// Handles errors
func Handle(err error, str string) {
	fmt.Printf("Error: %s\n", str)
	fmt.Println(err.Error())

}

func Log(info interface{}) {
	logger := log.New(os.Stderr).WithColor()
	logger.Info(info)
}

func JSON(obj interface{}) {
	logger := log.New(os.Stderr).WithColor()

	if bytes, err := json.MarshalIndent(obj, "  ", ""); err != nil {
		logger.Info(obj)
	} else {
		logger.Info(string(bytes))
	}
}

func Error(err error, info string) {
	logger := log.New(os.Stderr)
	logger.Info(info)
	logger.Error(err)
}
