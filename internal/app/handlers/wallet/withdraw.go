package wallet

import (
	"net/http"

	"github.com/labstack/echo/v4"
	pkgappctx "github.com/leeliwei930/walletassignment/internal/app/context"
	"github.com/leeliwei930/walletassignment/internal/app/models"
	"github.com/leeliwei930/walletassignment/internal/app/response"
)

type WithdrawalRequest struct {
	Amount int `json:"amount" validate:"required,min=1,max=1000000" localeKey:"wallet::withdraw::amount"`
}

type WithdrawalResponse struct {
	Wallet      models.WalletStatus      `json:"wallet"`
	Transaction models.WalletTransaction `json:"transaction"`
}

func (h *WalletHandler) Withdraw(c echo.Context) error {
	app := h.app
	responder := response.NewResponder(c, app)

	var withdrawalRequest WithdrawalRequest
	if err := c.Bind(&withdrawalRequest); err != nil {
		return responder.AbortIfIncorrectJsonPayload(c, err)
	}

	validator := app.GetValidator()
	if err := validator.Struct(withdrawalRequest); err != nil {
		return err
	}

	ctx := c.Request().Context()
	walletSvc := h.app.GetWalletService()
	appCtx, err := pkgappctx.GetApplicationContext(ctx)
	if err != nil {
		return err
	}

	userID := appCtx.GetAuthUserID()
	withdrawTrx, err := walletSvc.Withdraw(ctx, models.WalletWithdrawalParams{
		UserID: userID,
		Amount: withdrawalRequest.Amount,
	})

	if err != nil {
		return err
	}

	return responder.JSON(http.StatusOK, WithdrawalResponse{
		Wallet:      withdrawTrx.Wallet,
		Transaction: withdrawTrx.Transaction,
	})
}
