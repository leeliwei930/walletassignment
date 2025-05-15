package wallet

import "github.com/leeliwei930/walletassignment/internal/interfaces"

type WalletHandler struct {
	app interfaces.Application
}

func NewHandler(app interfaces.Application) *WalletHandler {
	return &WalletHandler{
		app: app,
	}
}
