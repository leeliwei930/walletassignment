package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/leeliwei930/walletassignment/ent"
)

type UserService interface {
	GetUserIDByPhone(ctx context.Context, phone string) (uuid.UUID, error)
	GetFullName(ctx context.Context, userRec *ent.User) string
	SetupUser(ctx context.Context, phone string, firstName string, lastName string) (*ent.User, error)
}
