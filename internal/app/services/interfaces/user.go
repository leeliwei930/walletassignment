package interfaces

import (
	"context"

	"github.com/leeliwei930/walletassignment/ent"
)

type UserService interface {
	GetUserByPhone(ctx context.Context, phone string) (*ent.User, error)
}
