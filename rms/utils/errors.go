package utils

import (
	"fmt"
	"log"
)

func Fatal(errMessage error, err error) {
	if err != nil {
		log.Fatal(errMessage)
		fmt.Println()
	}
}

// Returns true if error, else false
func Error(errMessage error, err error) bool {
	if err != nil {
		fmt.Println(errMessage)
		return true
	}

	return false
}
