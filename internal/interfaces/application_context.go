package interfaces

type ApplicationContext interface {
	GetAuthUserPhone() string
	GetLanguage() string

	SetLanguage(language string)
	SetAuthUserPhone(phone string)
}
