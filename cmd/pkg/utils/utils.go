package utils

import (
	"log"
)


func ErrorHandler(err error) {
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
	}
}


func GetSessionFolder(phone string) string {
	var out []rune
	for _, r := range phone {
		if r >= '0' && r <= '9' {
			out = append(out, r)
		}
	}
	return "phone-" + string(out)
}