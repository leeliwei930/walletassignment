package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/leeliwei930/walletassignment/constant"
	"github.com/leeliwei930/walletassignment/internal/app"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type ApplicationVersion struct {
	CommitSHA       string
	BuildReleaseTag string
}

func Start(appVersion *ApplicationVersion) {
	ec := echo.New()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	app, err := app.InitializeFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	app.SetupMiddlewares(ec)
	app.Routes(ec)

	// Start the server in a goroutine
	go func() {
		defer app.Close()
		log := app.GetLog()
		config := app.GetConfig()
		if config.Server == nil {
			log.Fatal("Server configuration is not set")
		}

		log.Info("App Environment",
			zap.String("environment", viper.GetString(constant.APP_ENV)),
		)
		log.Info("core services started on port", zap.Int("port", config.Server.Port))
		log.Info("Starting server...")
		log.Info("Server is running on", zap.String("host", config.Server.Host))

		err := ec.Start(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port))
		if err != nil {
			log.Fatal("Unable to start server")
		}
	}()
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := ec.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
