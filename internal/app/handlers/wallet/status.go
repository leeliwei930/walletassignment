package wallet

import (
	"net/http"

	"github.com/labstack/echo/v4"
	pkgappctx "github.com/leeliwei930/walletassignment/internal/app/context"
	"github.com/leeliwei930/walletassignment/internal/app/models"
	"github.com/leeliwei930/walletassignment/internal/app/response"
)

type StatusResponse struct {
	Wallet *models.WalletStatus `json:"wallet"`
}

func (h *WalletHandler) Status(ec echo.Context) error {
	ctx := ec.Request().Context()
	responder := response.NewResponder(ec, h.app)
	walletService := h.app.GetWalletService()

	appCtx, err := pkgappctx.GetApplicationContext(ctx)
	if err != nil {
		return responder.UnexpectedError(ec, err)
	}

	userID := appCtx.GetAuthUserID()
	params := models.WalletStatusParams{
		UserID: userID,
	}

	walletStatus, err := walletService.Status(ctx, params)
	if err != nil {
		return responder.UnexpectedError(ec, err)
	}

	return responder.JSON(http.StatusOK, StatusResponse{
		Wallet: walletStatus,
	})
}
