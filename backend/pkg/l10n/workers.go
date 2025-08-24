package l10n

import "errors"

const (
	GenericWorker  = "Ошибка рабочего"
	WorkerNotFound = "Рабочий не найден"
	InvalidRole    = "Некорректная роль"
	WorkerExists   = "Рабочий уже существует"
)

var (
	GenericWorkerError  = errors.New(GenericWorker)
	WorkerNotFoundError = errors.New(WorkerNotFound)
	InvalidRoleError    = errors.New(InvalidRole)
	WorkerExistsError   = errors.New(WorkerExists)
)
