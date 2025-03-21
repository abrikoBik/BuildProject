package util

import (
	"errors"
	"log"
)

var (
	WrongPasswordError = errors.New("wrong password")
	InvalidEmailError  = errors.New("invalid email")
)

func Check(errs ...error) {
	joined := errors.Join(errs...)
	if joined != nil {
		log.Fatalln(joined)
	}
}
