package app

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/leeliwei930/walletassignment/config"
	"github.com/leeliwei930/walletassignment/ent"
	"github.com/leeliwei930/walletassignment/internal/interfaces"
	"go.uber.org/zap"
)

type application struct {
	ent        *ent.Client
	db         *sql.DB
	config     *config.Config
	dbMigrator interfaces.DBMigrator
	log        *zap.Logger
	validator  *validator.Validate
	locale     interfaces.Locale
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

func (app *application) GetLog() *zap.Logger {
	return app.log
}

func (app *application) GetConfig() *config.Config {
	return app.config
}

func (app *application) GetDBMigrator() interfaces.DBMigrator {
	return app.dbMigrator
}

func (app *application) GetValidator() *validator.Validate {
	return app.validator
}

func (app *application) GetLocale() interfaces.Locale {
	return app.locale
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
