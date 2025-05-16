package app

import (
	svcinterfaces "github.com/leeliwei930/walletassignment/internal/app/services/interfaces"
	pkgusersvc "github.com/leeliwei930/walletassignment/internal/app/services/user"
	pkgwalletsvc "github.com/leeliwei930/walletassignment/internal/app/services/wallet"
)

func (a *application) GetUserService() svcinterfaces.UserService {
	return a.userService
}

func (a *application) GetWalletService() svcinterfaces.WalletService {
	return a.walletService
}

func (app *application) InitServices() error {
	app.userService = pkgusersvc.New(app)
	app.walletService = pkgwalletsvc.New(app)
	return nil
}
