package interfaces

import (
	"context"

	"github.com/google/uuid"
)

type UserService interface {
	GetUserIDByPhone(ctx context.Context, phone string) (uuid.UUID, error)
}
