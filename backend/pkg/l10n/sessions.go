package l10n

import "errors"

const (
	GenericSession    = "Ошибка сессии"
	TokenNotFound     = "Токен не найден"
	InvalidToken      = "Некорректный токен"
	InvalidAuthHeader = "Некорректный заголовок аутентификации"
	Login             = "Ошибка входа"
	Logout            = "Ошибка выхода"
)

var (
	GenericSessionError    = errors.New(GenericSession)
	TokenNotFoundError     = errors.New(TokenNotFound)
	InvalidTokenError      = errors.New(InvalidToken)
	InvalidAuthHeaderError = errors.New(InvalidAuthHeader)
	LoginError             = errors.New(Login)
	LogoutError            = errors.New(Logout)
)
