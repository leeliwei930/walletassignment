package database

import (
	"database/sql"

	entSQL "entgo.io/ent/dialect/sql"
	"github.com/leeliwei930/walletassignment/config"
	"github.com/leeliwei930/walletassignment/ent"
)

const (
	DBConnectionSetupErr = "unable setup database connection to %s"
)

func BuildEntClient(drv *entSQL.Driver) *ent.Client {
	entClient := ent.NewClient(ent.Driver(drv))

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

	drv := entSQL.OpenDB(config.Connection.Driver(), db)

	return drv, nil
}
