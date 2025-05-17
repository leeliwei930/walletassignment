package app

import (
	"github.com/labstack/echo/v4"
	"github.com/leeliwei930/walletassignment/internal/app/handlers"
	"github.com/leeliwei930/walletassignment/internal/app/handlers/wallet"
	"github.com/leeliwei930/walletassignment/internal/app/middleware"
)

func (app *application) Routes(ec *echo.Echo) *echo.Echo {

	ec.HTTPErrorHandler = handlers.ApplicationErrorHandler(app)

	walletHandler := wallet.NewHandler(app)
	handler := handlers.NewHandler(app)

	api := ec.Group("/api")

	v1 := api.Group("/v1")
	v1.GET("/health", handler.Health)

	wallet := v1.Group("/wallet")

	wallet.Use(middleware.RequireAuth(app))
	wallet.GET("/status", walletHandler.Status)
	wallet.POST("/deposit", walletHandler.Deposit)
	wallet.POST("/withdraw", walletHandler.Withdraw)
	wallet.POST("/transfer", walletHandler.Transfer)
	wallet.GET("/transactions", walletHandler.GetTransactions)

	return ec
}
