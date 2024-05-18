package utils

import (
	"log"
)


func ErrorHandler(err error) {
	if err != nil {
		log.Panicf("ERROR: %s", err.Error())
	}
}