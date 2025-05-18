package database

import (
	"database/sql"
	"testing"

	entSQL "entgo.io/ent/dialect/sql"
	"github.com/leeliwei930/walletassignment/config"
	"github.com/leeliwei930/walletassignment/ent"
	"github.com/leeliwei930/walletassignment/ent/enttest"
)

const (
	DBConnectionSetupErr = "unable setup database connection to %s"
)

func BuildEntClient(drv *entSQL.Driver) *ent.Client {
	entClient := ent.NewClient(ent.Driver(drv))

	return entClient
}

func BuildEntTestClient(t *testing.T, drv *entSQL.Driver) *ent.Client {
	entClient := enttest.NewClient(t, enttest.WithOptions(
		ent.Driver(drv),
	))

	return entClient
}

func BuildSQLDriver(config *config.DBConnectionConfig) (*entSQL.Driver, error) {
	db, err := sql.Open(config.Connection.Driver(), config.Connection.DSN())
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(config.MaxIdleConnections)
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetConnMaxLifetime(config.ConnectionMaxLifeTime)

	drv := entSQL.OpenDB(config.Connection.EntDialect(), db)

	return drv, nil
}
