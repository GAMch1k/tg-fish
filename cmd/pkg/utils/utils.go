package utils

import (
	"log"
)


func ErrorHandler(err error) {
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
	}
}
