package app

import (
	"context"
	"errors"

	"github.com/leeliwei930/walletassignment/database"
	"github.com/leeliwei930/walletassignment/database/migrator"
)

func (app *application) InitEntClient() error {
	config := app.GetConfig().DBConfig
	if config == nil {
		return errors.New("DBConfig is not set in configuration")
	}
	drv, err := database.BuildSQLDriver(config)
	if err != nil {
		return err
	}

	app.db = drv.DB()

	entClient := database.BuildEntClient(drv)
	app.ent = entClient
	return nil
}

func (app *application) WrapRefreshDatabaseTransaction(ctx context.Context, fn func()) error {
	config := app.GetConfig().DBConfig
	if config == nil {
		return errors.New("DBConfig is not set in configuration")
	}
	drv, err := database.BuildSQLDriver(config)
	if err != nil {
		return err
	}

	app.db = drv.DB()

	entClient := database.BuildEntClient(drv)

	entTx, err := entClient.Tx(ctx)
	if err != nil {
		return err
	}

	app.ent = entTx.Client()

	fn()

	entTx.Rollback()
	app.InitEntClient()

	return nil
}

func (app *application) SetupDBMigrator() {
	atlasMigrator := migrator.NewAtlasDBMigrator(app)
	app.dbMigrator = atlasMigrator
}
