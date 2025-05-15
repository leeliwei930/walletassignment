package wallet

import (
	"github.com/leeliwei930/walletassignment/internal/interfaces"

	walletinterfaces "github.com/leeliwei930/walletassignment/internal/app/services/interfaces"
)

type walletService struct {
	app interfaces.Application
}

func NewWalletService(app interfaces.Application) walletinterfaces.WalletService {
	return &walletService{app: app}
}
