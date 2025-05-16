package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/leeliwei930/walletassignment/ent/user"
)

func (s *userService) GetUserIDByPhone(ctx context.Context, phone string) (uuid.UUID, error) {
	client := s.app.GetEnt()
	userClient := client.User

	user, err := userClient.Query().Where(
		user.PhoneNumberEQ(phone),
	).First(ctx)

	if err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}
