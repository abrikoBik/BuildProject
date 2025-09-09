package l10n

import "errors"

const (
	GenericProject  = "Ошибка проекта"
	ProjectNotFound = "Проект не найден"
)

var (
	GenericProjectError  = errors.New(GenericProject)
	ProjectNotFoundError = errors.New(ProjectNotFound)
)
