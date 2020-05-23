package loglib

import "log"

func CheckFatall(err error, text string) {
	if err != nil {
		log.Fatal(text, err)
	}
}
