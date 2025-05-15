package app

import (
	"github.com/labstack/echo/v4"
)

func (app *application) Routes(ec *echo.Echo) *echo.Echo {

	api := ec.Group("/api")
	v1 := api.Group("/v1")

	v1.Group("/wallet")
	return ec
}
