package l10n

import "errors"

const (
	GenericStage  = "Ошибка стадии"
	StageNotFound = "Стадия не найдена"

	GenericStageGroup  = "Ошибка группы стадий"
	StageGroupNotFound = "Группа стадий не найдена"
)

var (
	GenericStageError  = errors.New(GenericStage)
	StageNotFoundError = errors.New(StageNotFound)

	GenericStageGroupError  = errors.New(GenericStageGroup)
	StageGroupNotFoundError = errors.New(StageGroupNotFound)
)
