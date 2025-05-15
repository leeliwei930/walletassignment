package app

import (
	"github.com/labstack/echo/v4"
	"github.com/leeliwei930/walletassignment/internal/app/handlers"
)

func (app *application) Routes(ec *echo.Echo) *echo.Echo {

	api := ec.Group("/api")
	v1 := api.Group("/v1")

	wallet := v1.Group("/wallet")

	walletHandler := handlers.NewWalletHandler(app)

	wallet.GET("/health", walletHandler.Health)

	return ec
}
