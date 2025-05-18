package wallet

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	pkgappcontext "github.com/leeliwei930/walletassignment/internal/app/context"
	"github.com/leeliwei930/walletassignment/internal/app/models"
)

type WalletTransactionsResponse struct {
	Data       []*models.WalletTransaction `json:"data"`
	Pagination *models.Pagination          `json:"pagination"`
}

type WalletTransactionsRequest struct {
	Limit int `json:"limit" validate:"omitempty,min=1,max=100" localeKey:"wallet::transactions::limit"`
	Page  int `json:"page" validate:"omitempty,min=1" localeKey:"wallet::transactions::page"`
}

func (h *WalletHandler) GetTransactions(c echo.Context) error {
	app := h.app
	ctx := c.Request().Context()
	appCtx, err := pkgappcontext.GetApplicationContext(ctx)
	if err != nil {
		return err
	}

	authUserID := appCtx.GetAuthUserID()

	limit := 0
	page := 0

	if c.QueryParams().Has("limit") {
		limitStr := c.QueryParam("limit")
		limit, _ = strconv.Atoi(limitStr)
	}

	if c.QueryParams().Has("page") {
		pageStr := c.QueryParam("page")
		page, _ = strconv.Atoi(pageStr)
	}

	request := WalletTransactionsRequest{
		Limit: limit,
		Page:  page,
	}
	validator := app.GetValidator()
	if err := validator.Struct(request); err != nil {
		return err
	}

	wallestSvc := app.GetWalletService()
	walletTransactions, err := wallestSvc.Transactions(ctx, models.WalletTransactionsParams{
		UserID: authUserID,
		Limit:  limit,
		Page:   page,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, WalletTransactionsResponse{
		Data:       walletTransactions.Transactions,
		Pagination: walletTransactions.Pagination,
	})
}
