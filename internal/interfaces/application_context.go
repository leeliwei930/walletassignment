package interfaces

import (
	"github.com/google/uuid"
)

type ApplicationContext interface {
	GetAuthUserID() uuid.UUID
	GetLanguage() string

	SetLanguage(language string)
	SetAuthUserID(userID uuid.UUID)
}
