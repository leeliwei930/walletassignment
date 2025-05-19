package app

import (
	"code.cloudfoundry.org/clock"
	svcinterfaces "github.com/leeliwei930/walletassignment/internal/app/services/interfaces"
)

func (a *application) GetClock() clock.Clock {
	return a.clock
}

func (a *application) SetClock(c clock.Clock) {
	a.clock = c
}

func (a *application) SetWalletService(ws svcinterfaces.WalletService) {
	a.walletService = ws
}

func (a *application) SetUserService(us svcinterfaces.UserService) {
	a.userService = us
}
