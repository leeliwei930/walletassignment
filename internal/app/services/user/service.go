package user

import (
	"github.com/leeliwei930/walletassignment/internal/app/services/interfaces"
	appinterfaces "github.com/leeliwei930/walletassignment/internal/interfaces"
)

type userService struct {
	app appinterfaces.Application
}

func New(app appinterfaces.Application) interfaces.UserService {
	return &userService{app: app}
}
