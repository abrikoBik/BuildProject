package l10n

import "errors"

const (
	GenericEstimate  = "Ошибка сметы"
	EstimateNotFound = "Смета не найдена"

	GenericEstimateGroup  = "Ошибка группы смет"
	EstimateGroupNotFound = "Группа смет не найдена"
)

var (
	GenericEstimateError  = errors.New(GenericEstimate)
	EstimateNotFoundError = errors.New(EstimateNotFound)

	GenericEstimateGroupError  = errors.New(GenericEstimateGroup)
	EstimateGroupNotFoundError = errors.New(EstimateGroupNotFound)
)
