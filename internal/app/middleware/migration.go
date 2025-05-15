package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/leeliwei930/walletassignment/internal/app/response"
	"github.com/leeliwei930/walletassignment/internal/errors"
	"github.com/leeliwei930/walletassignment/internal/interfaces"
	pkgerrors "github.com/pkg/errors"
)

func MigrationPrecheck(app interfaces.Application) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			dbMigrator := app.GetDBMigrator()
			err := dbMigrator.PendingMigrationCheck(ctx)
			responder := response.NewResponder(c, app)

			if err != nil {
				err = errors.NewStandardError(
					"ERR_PENDING_DB_MIGRATION_500",
					"Unable to start the services due to pending database migration",
					pkgerrors.New("Unable to start the services due to pending database migration"),
				)
				return responder.UnexpectedError(c, err)
			}

			return next(c)
		}
	}
}
