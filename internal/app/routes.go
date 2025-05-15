package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *application) Routes(ec *echo.Echo) *echo.Echo {

	api := ec.Group("/api")
	v1 := api.Group("/v1")

	wallet := v1.Group("/wallet")

	wallet.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	return ec
}
