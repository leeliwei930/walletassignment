package app

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/leeliwei930/walletassignment/constant"
	pkgappmiddleware "github.com/leeliwei930/walletassignment/internal/app/middleware"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func (app *application) SetupMiddlewares(ec *echo.Echo) *echo.Echo {
	ec.IPExtractor = echo.ExtractIPFromXFFHeader()

	ec.Pre(middleware.RemoveTrailingSlash())
	ec.Use(pkgappmiddleware.MigrationPrecheck(app))
	ec.Use(middleware.RecoverWithConfig(
		middleware.RecoverConfig{
			StackSize: 16 << 10, // 1 KB
			LogErrorFunc: func(ctx echo.Context, err error, stack []byte) error {
				app.GetLog().Error(err.Error(),
					zap.Bool("isPanic", true),
					zap.Error(err),
					zap.String("stack", string(stack)),
				)
				return err
			},
		},
	))
	ec.Use(pkgappmiddleware.LanguageMiddleware(app))
	ec.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogRequestID: true,
		LogRemoteIP:  true,
		LogLatency:   true,
		LogMethod:    true,
		LogURIPath:   true,
		LogRoutePath: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			app.GetLog().Info("wallet-api-request",
				zap.String("request_id", v.RequestID),
				zap.String("remote_ip", v.RemoteIP),
				zap.String("method", v.Method),
				zap.String("latency", v.Latency.String()),
				zap.String("path", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	}))
	ec.Use(middleware.RequestID())
	ec.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	ec.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     strings.Split(viper.GetString(constant.CORS_ALLOWED_ORIGINS), ","),
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	return ec
}
