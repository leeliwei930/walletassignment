package wallet

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/leeliwei930/walletassignment/ent"
	"github.com/leeliwei930/walletassignment/internal/app/models"
	"github.com/leeliwei930/walletassignment/internal/app/response"
	"github.com/leeliwei930/walletassignment/internal/errors"
)

type TransferRequest struct {
	Amount               int    `json:"amount" validate:"required,min=1,max=1000000" localeKey:"wallet::transfer::amount"`
	RecipientPhoneNumber string `json:"recipient_phone_number" validate:"required" localeKey:"wallet::transfer::recipient_phone_number"`
}

type TransferResponse struct {
	Wallet      models.WalletStatus      `json:"wallet"`
	Transaction models.WalletTransaction `json:"transaction"`
}

func (h *WalletHandler) Transfer(ec echo.Context) error {

	app := h.app
	responder := response.NewResponder(ec, app)
	ctx := ec.Request().Context()
	var transferRequest TransferRequest
	if err := ec.Bind(&transferRequest); err != nil {
		return responder.AbortIfIncorrectJsonPayload(ec, err)
	}

	walletSvc := app.GetWalletService()
	userSvc := app.GetUserService()

	recipientUserID, err := userSvc.GetUserIDByPhone(ctx, transferRequest.RecipientPhoneNumber)
	if err != nil && ent.IsNotFound(err) {
		return errors.InvalidTransferRecipientPhoneNumberErr
	} else if err != nil {
		return err
	}

	walletTransfer, err := walletSvc.Transfer(ctx, models.WalletTransferParams{
		RecipientUserID: recipientUserID,
		Amount:          transferRequest.Amount,
	})
	if err != nil {
		return err
	}

	return responder.JSON(http.StatusOK, TransferResponse{
		Wallet:      walletTransfer.Wallet,
		Transaction: walletTransfer.Transaction,
	})
}
