package l10n

import "errors"

const (
	GenericUser      = "Ошибка пользователя"
	WrongPassword    = "Неверный пароль"
	InsecurePassword = "Минимальная длина пароля - 8 символов"
	FieldTaken       = "Имя пользователя, электронная почта или номер телефона уже заняты"
	InvalidEmail     = "Некорректная почта"
	UserNotFound     = "Пользователь не найден"
	RegisterError    = "Ошибка регистрации"
	EmptyPassword    = "Пароль не может быть пустым"
)

var (
	GenericUserError      = errors.New(GenericUser)
	WrongPasswordError    = errors.New(WrongPassword)
	InsecurePasswordError = errors.New(InsecurePassword)
	FieldTakenError       = errors.New(FieldTaken)
	InvalidEmailError     = errors.New(InvalidEmail)
	UserNotFoundError     = errors.New(UserNotFound)
	RegisterErrorError    = errors.New(RegisterError)
	EmptyPasswordError    = errors.New(EmptyPassword)
)
