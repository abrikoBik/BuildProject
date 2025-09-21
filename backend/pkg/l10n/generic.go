package l10n

import "errors"

const (
	OutOfBounds        = "Индекс вне диапазона"
	MissingID          = "Отсутствует ID"
	InvalidID          = "Некорректный ID"
	ContextID          = "Ошибка получения ID"
	BadRequest         = "Плохой JSON-запрос"
	ResponseGeneration = "Ошибка при создании JSON-ответа"
)

var (
	OutOfBoundsError        = errors.New(OutOfBounds)
	MissingIDError          = errors.New(MissingID)
	InvalidIDError          = errors.New(InvalidID)
	ContextIDError          = errors.New(ContextID)
	BadRequestError         = errors.New(BadRequest)
	ResponseGenerationError = errors.New(ResponseGeneration)
)
