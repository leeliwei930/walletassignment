package wallet

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/leeliwei930/walletassignment/internal/app/models"
	"github.com/leeliwei930/walletassignment/internal/app/response"
)

type DepositRequest struct {
	// Minimum deposit amount is 1 USD and maximum is 10000 USD, value represent in cents
	Amount int `json:"amount" validate:"required,min=100,max=1000000" localeKey:"wallet::deposit::amount"`
}

type DepositResponse struct {
	Success     bool                     `json:"success"`
	Wallet      models.WalletStatus      `json:"wallet"`
	Transaction models.WalletTransaction `json:"transaction"`
}

func (h *WalletHandler) Deposit(ec echo.Context) error {

	app := h.app
	responder := response.NewResponder(ec, app)

	var depositRequest DepositRequest
	if err := ec.Bind(&depositRequest); err != nil {
		return responder.AbortIfIncorrectJsonPayload(ec, err)
	}

	validator := app.GetValidator()
	if err := validator.Struct(depositRequest); err != nil {
		return err
	}

	ctx := ec.Request().Context()
	walletSvc := h.app.GetWalletService()
	depositTrx, err := walletSvc.Deposit(ctx, models.WalletDepositParams{
		Amount: depositRequest.Amount,
	})
	if err != nil {
		return err
	}

	return responder.JSON(http.StatusOK, DepositResponse{
		Success:     true,
		Wallet:      depositTrx.Wallet,
		Transaction: depositTrx.Transaction,
	})
}
