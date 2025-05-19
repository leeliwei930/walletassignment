package app

import "code.cloudfoundry.org/clock"

func (a *application) GetClock() clock.Clock {
	return a.clock
}

func (a *application) SetClock(c clock.Clock) {
	a.clock = c
}
