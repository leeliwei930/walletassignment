package middleware

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/leeliwei930/walletassignment/constant"
	pkgappcontext "github.com/leeliwei930/walletassignment/internal/app/context"
	"github.com/leeliwei930/walletassignment/internal/interfaces"
)

func LanguageMiddleware(app interfaces.Application) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			appCtx := pkgappcontext.New()

			req := c.Request()
			reqHeaders := req.Header
			ctx := req.Context()
			ctx = context.WithValue(ctx, constant.ApplicationContextKey, appCtx)

			acceptLang := reqHeaders.Get("Accept-Language")
			appCtx.SetLanguage(acceptLang)

			req = req.WithContext(ctx)
			c.SetRequest(req)

			return next(c)
		}
	}
}
