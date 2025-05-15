package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/leeliwei930/walletassignment/ent"
	"github.com/leeliwei930/walletassignment/ent/user"
)

func (s *userService) GetUserByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	client := s.app.GetEnt()

	userClient := client.User
	return userClient.Query().Where(
		user.IDEQ(id),
	).First(ctx)
}
