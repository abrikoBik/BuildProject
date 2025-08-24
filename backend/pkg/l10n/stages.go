package l10n

import "errors"

const (
	GenericStage  = "Ошибка стадии"
	StageNotFound = "Стадия не найдена"
)

var (
	GenericStageError  = errors.New(GenericStage)
	StageNotFoundError = errors.New(StageNotFound)
)
