package app

import (
	"github.com/leeliwei930/walletassignment/config"
	"github.com/leeliwei930/walletassignment/ent"
	"github.com/spf13/viper"
)

type InitializeOpts = func(*application) *application

func WithConfig(config *config.Config) func(*application) *application {
	return func(app *application) *application {
		app.config = config
		return app
	}
}

func WithEntClient(entClient *ent.Client) func(*application) *application {
	return func(app *application) *application {
		app.ent = entClient
		return app
	}
}

func InitializeFromEnv(opts ...InitializeOpts) (*application, error) {

	viper.AutomaticEnv()
	app := New()
	var err error

	for _, opt := range opts {
		app = opt(app)
	}

	err = initializeAll(
		app.initializeConfigIfNeeded(),
		app.InitEntClient(),
	)
	if err != nil {
		return nil, err
	}

	app.SetupDBMigrator()

	return app, nil
}

func (app *application) initializeConfigIfNeeded() error {
	if app.config == nil {
		return app.InitConfig()
	}
	return nil
}

type initializers = error

func initializeAll(opts ...initializers) error {
	for _, err := range opts {
		if err != nil {
			return err
		}
	}
	return nil
}
