package util

import (
	"errors"
	"log"
	"mp1/pkg/l10n"
)

var (
	InsecurePasswordError = errors.New(l10n.InsecurePassword)
	WrongPasswordError    = errors.New(l10n.WrongPassword)
	InvalidEmailError     = errors.New(l10n.InvalidEmail)
	InvalidRoleError      = errors.New(l10n.InvalidRole)
	InvalidIDError        = errors.New(l10n.InvalidID)
	InvalidTokenError     = errors.New(l10n.InvalidToken)
)

func Check(errs ...error) {
	joined := errors.Join(errs...)
	if joined != nil {
		log.Fatalln(joined)
	}
}
