package app

import (
	"context"
	"errors"
	"testing"

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

func (app *application) UseRefreshDB(t *testing.T, fn func()) error {
	config := app.GetConfig().DBConfig
	if config == nil {
		return errors.New("DBConfig is not set in configuration")
	}
	drv, err := database.BuildSQLDriver(config)
	if err != nil {
		return err
	}

	app.db = drv.DB()
	app.ent = database.BuildEntTestClient(t, drv)
	defer func() {
		_, _ = app.ent.Ledger.Delete().Exec(context.Background())
		_, _ = app.ent.Wallet.Delete().Exec(context.Background())
		_, _ = app.ent.User.Delete().Exec(context.Background())

		app.ent.Close()
	}()

	fn()

	return nil
}

func (app *application) SetupDBMigrator() {
	atlasMigrator := migrator.NewAtlasDBMigrator(app)
	app.dbMigrator = atlasMigrator
}
