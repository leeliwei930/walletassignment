package user

import (
	"context"
	"fmt"

	"github.com/leeliwei930/walletassignment/ent"
)

func (s *userService) GetFullName(ctx context.Context, userRec *ent.User) string {
	return fmt.Sprintf("%s %s", userRec.FirstName, userRec.LastName)
}
