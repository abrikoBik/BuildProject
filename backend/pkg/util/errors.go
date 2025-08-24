package util

import (
	"errors"
	"log"
)

func Check(errs ...error) {
	joined := errors.Join(errs...)
	if joined != nil {
		log.Fatalln(joined)
	}
}
