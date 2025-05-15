package middleware

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/leeliwei930/walletassignment/internal/app/response"
	"github.com/leeliwei930/walletassignment/internal/interfaces"
)

func RequireAuth(app interfaces.Application) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			responder := response.NewResponder(c, app)

			headers := c.Request().Header
			userPhone := headers.Get("X-USER-PHONE")

			if len(userPhone) == 0 {
				return responder.UnauthorizedError(c, errors.New("X-USER-PHONE is required to present in header"))
			}

			return next(c)
		}
	}
}
