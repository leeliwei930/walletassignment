package interfaces

import (
	"context"

	"ariga.io/atlas-go-sdk/atlasexec"
)

// DBMigrator defines the interface for database migration operations
type DBMigrator interface {
	// GetMigrationStatus retrieves the current migration status
	GetMigrationStatus(ctx context.Context) (*atlasexec.MigrateStatus, error)

	// ApplyMigration applies pending migrations
	ApplyMigration(ctx context.Context, targetVersion string, dryRun bool) (*atlasexec.MigrateApply, error)

	// RollbackMigration rolls back migrations to a specific version
	RollbackMigration(ctx context.Context, targetVersion string) (*atlasexec.MigrateDown, error)

	// RehashMigration rehashes migration files
	RehashMigration(ctx context.Context) error

	// PendingMigrationCheck checks if there are pending migrations
	PendingMigrationCheck(ctx context.Context) error

	// Close closes the DBMigrator
	Close() error
}
