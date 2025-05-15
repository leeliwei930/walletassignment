package interfaces

import (
	"context"
	"database/sql"

	"github.com/leeliwei930/walletassignment/config"
	"github.com/leeliwei930/walletassignment/ent"
	"go.uber.org/zap"
)

type Application interface {
	GetEnt() *ent.Client
	GetDB() *sql.DB
	GetConfig() *config.Config
	GetDBMigrator() DBMigrator
	GetLog() *zap.Logger
	Close()
	WrapRefreshDatabaseTransaction(ctx context.Context, fn func()) error
}
