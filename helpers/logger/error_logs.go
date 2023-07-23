package logger

import (
	"fmt"
)

// Handles errors
func Handle(err error, str string) {
	fmt.Printf("Error: %s\n", str)
	fmt.Println(err.Error())
}
