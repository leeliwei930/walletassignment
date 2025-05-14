package app

import (
	"database/sql"

	"github.com/leeliwei930/walletassignment/config"
	"github.com/leeliwei930/walletassignment/ent"
	"github.com/leeliwei930/walletassignment/internal/interfaces"
)

type application struct {
	ent        *ent.Client
	db         *sql.DB
	config     *config.Config
	dbMigrator interfaces.DBMigrator
}

func New() *application {
	return &application{}
}

func (app *application) GetEnt() *ent.Client {
	return app.ent
}

func (app *application) GetDB() *sql.DB {
	return app.db
}

func (app *application) GetConfig() *config.Config {
	return app.config
}

func (app *application) GetDBMigrator() interfaces.DBMigrator {
	return app.dbMigrator
}

func (app *application) Close() {
	if app.ent != nil {
		app.ent.Close()
	}
	if app.db != nil {
		app.db.Close()
	}

	if app.dbMigrator != nil {
		app.dbMigrator.Close()
	}
}
