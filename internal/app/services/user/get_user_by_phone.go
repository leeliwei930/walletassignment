package user

import (
	"context"

	"github.com/leeliwei930/walletassignment/ent"
	"github.com/leeliwei930/walletassignment/ent/user"
)

func (s *userService) GetUserByPhone(ctx context.Context, phone string) (*ent.User, error) {

	client := s.app.GetEnt()

	userClient := client.User
	return userClient.Query().Where(
		user.PhoneNumberEQ(phone),
	).First(ctx)
}
