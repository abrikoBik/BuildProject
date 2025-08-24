package l10n

import "errors"

const (
	OutOfBounds = "Индекс вне диапазона"
	MissingID   = "Отсутствует ID"
	InvalidID   = "Некорректный ID"
)

var (
	OutOfBoundsError = errors.New(OutOfBounds)
	MissingIDError   = errors.New(MissingID)
	InvalidIDError   = errors.New(InvalidID)
)
