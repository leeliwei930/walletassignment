package interfaces

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/leeliwei930/walletassignment/config"
	"github.com/leeliwei930/walletassignment/ent"
	svcinterfaces "github.com/leeliwei930/walletassignment/internal/app/services/interfaces"
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
	GetLocale() Locale
	GetUserService() svcinterfaces.UserService
	GetWalletService() svcinterfaces.WalletService
	GetValidator() *validator.Validate
}
