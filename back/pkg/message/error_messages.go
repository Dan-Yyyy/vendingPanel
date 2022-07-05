package message

const (
	EmptyAuthHeader   = "пустой заголовок авторизации"
	InvalidAuthHeader = "неверный формат заголовка авторизации"

	ReadEnvError = "ошибка считывания env файла"

	DBConnectionError = "не удалось подключиться к БД"

	StartHTTPServerError = "возникла ошибка запуска https сервера"

	InvalidSigningMethod = "невалидный метод подписи"
	UnknownTokenClaims   = "тип token claims должны быть типа *tokenClaims"
	InvalidTypeOfUserId  = "невалидный тип айди пользователя"

	SeedingError = "ошибка сидинга"

	CreatUserError   = "не удалось создать пользователя"
	UserNotFound     = "пользователь с такими данными не найден"
	UserAlreadyExist = "пользователь с такой почтой уже существует"
)
