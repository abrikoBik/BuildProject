package util

import "log"

func Check(errors ...error) {
	for _, err := range errors {
		if err != nil {
			log.Fatalln(err)
		}
	}
}
