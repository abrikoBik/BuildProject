package l10n

const (
	// Invalid request errors
	InvalidID         = "Некорректный ID"
	InvalidAuthHeader = "Некорректный заголовок аутентификации"
	InvalidToken      = "Некорректный токен"
	InvalidEmail      = "Некорректная почта"
	InvalidRole       = "Некорректная роль"
	WrongPassword     = "Неверный пароль"
	InsecurePassword  = "Минимальная длина пароля - 8 символов"
	NameOrEmailTaken  = "Имя пользователя или почта уже занята"

	// Missing data errors
	MissingID     = "Отсутствует ID"
	EmptyPassword = "Пароль не может быть пустым"

	// "X not found" errors
	UserNotFound  = "Пользователь не найден"
	TokenNotFound = "Токен не найден"

	// Generic errors
	UserRegisterError = "Ошибка регистрации"
	LogoutError       = "Не удалось выйти"
)
