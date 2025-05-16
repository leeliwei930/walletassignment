package middleware

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/leeliwei930/walletassignment/constant"
	"github.com/leeliwei930/walletassignment/ent"
	pkgappcontext "github.com/leeliwei930/walletassignment/internal/app/context"
	"github.com/leeliwei930/walletassignment/internal/app/response"
	"github.com/leeliwei930/walletassignment/internal/interfaces"
)

func RequireAuth(app interfaces.Application) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			responder := response.NewResponder(c, app)
			req := c.Request()
			ctx := req.Context()
			headers := req.Header

			appCtx, err := pkgappcontext.GetApplicationContext(ctx)
			if err != nil {
				return responder.UnexpectedError(c, err)
			}

			userPhone := headers.Get("X-USER-PHONE")
			if len(userPhone) == 0 {
				return responder.UnauthorizedError(c)
			}

			userSvc := app.GetUserService()
			userID, err := userSvc.GetUserIDByPhone(ctx, userPhone)

			switch {
			case userID == uuid.Nil || ent.IsNotFound(err):
				return responder.UnauthorizedError(c)
			case err != nil:
				return responder.UnexpectedError(c, err)
			}

			appCtx.SetAuthUserID(userID)

			ctx = context.WithValue(ctx, constant.ApplicationContextKey, appCtx)
			req = req.WithContext(ctx)
			c.SetRequest(req)

			return next(c)
		}
	}
}
