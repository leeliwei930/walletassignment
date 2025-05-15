package handlers

import "github.com/leeliwei930/walletassignment/internal/interfaces"

type WalletHandler struct {
	app interfaces.Application
}

func NewWalletHandler(app interfaces.Application) *WalletHandler {
	return &WalletHandler{
		app: app,
	}
}
