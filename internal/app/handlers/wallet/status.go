package wallet

import (
	"net/http"

	"github.com/labstack/echo/v4"
	pkgappcontext "github.com/leeliwei930/walletassignment/internal/app/context"
)

func (h *WalletHandler) Status(c echo.Context) error {

	app := h.app
	req := c.Request()
	trans := app.GetLocale().GetTranslatorFromRequest(req)

	msg, _ := trans.T("errors::wallet::deposit::minimum_deposit_amount_required", "MYR 100")

	appCtx, err := pkgappcontext.GetApplicationContext(req.Context())
	if err != nil {
		panic(err)
	}

	phoneNo := appCtx.GetAuthUserPhone()
	return c.JSON(http.StatusOK, map[string]string{
		"message": msg,
		"phone":   phoneNo,
	})
}
