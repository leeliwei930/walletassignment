package migrator

import (
	"context"
	"fmt"
	"log"
	"os"

	"ariga.io/atlas-go-sdk/atlasexec"

	"github.com/leeliwei930/walletassignment/internal/errors"
	"github.com/leeliwei930/walletassignment/internal/interfaces"
	pkgerrors "github.com/pkg/errors"
)

type AtlasDBMigrator struct {
	app          interfaces.Application
	atlasClient  *atlasexec.Client
	migrationDir *atlasexec.WorkingDir
}

func PendingMigrationErrors(pendingMigrationsFile []atlasexec.File) error {
	errorMessage := "Pending migrations found. Please run the migration before starting the server \n"
	errorMessage += "List of migrations pending to be applied: \n"
	for idx, file := range pendingMigrationsFile {
		errorMessage += fmt.Sprintf("[%d] Version: %s, Name: %s \n", idx+1, file.Version, file.Name)
	}
	migrationError := pkgerrors.New(errorMessage)
	return errors.UnexpectedError(migrationError)
}

func NewAtlasDBMigrator(app interfaces.Application) interfaces.DBMigrator {

	migrationDir, err := atlasexec.NewWorkingDir(
		atlasexec.WithMigrations(
			os.DirFS("./database/migrations"),
		),
	)

	if err != nil {
		log.Fatal(errors.UnexpectedError(pkgerrors.Wrap(err, "Atlas database migration directory error")).Error())
		return nil
	}

	client, err := atlasexec.NewClient(migrationDir.Path(), "atlas")
	if err != nil {
		log.Fatal(errors.UnexpectedError(pkgerrors.Wrap(err, "Atlas database migration client setup error")).Error())
		return nil
	}

	return &AtlasDBMigrator{
		app:          app,
		atlasClient:  client,
		migrationDir: migrationDir,
	}
}

func (dbm *AtlasDBMigrator) GetMigrationStatus(ctx context.Context) (*atlasexec.MigrateStatus, error) {
	client := dbm.atlasClient
	config := dbm.app.GetConfig()
	dbConn := config.DBConfig.Connection

	return client.MigrateStatus(ctx, &atlasexec.MigrateStatusParams{
		URL: dbConn.AtlasDSN(),
	})
}

func (dbm *AtlasDBMigrator) ApplyMigration(ctx context.Context, targetVersion string, dryRun bool) (*atlasexec.MigrateApply, error) {
	client := dbm.atlasClient
	config := dbm.app.GetConfig()
	dbConn := config.DBConfig.Connection

	params := &atlasexec.MigrateApplyParams{
		URL: dbConn.AtlasDSN(),
	}

	if targetVersion != "" {
		params.BaselineVersion = targetVersion
	}

	if dryRun {
		params.DryRun = true
	}
	return client.MigrateApply(ctx, params)
}

func (dbm *AtlasDBMigrator) RollbackMigration(ctx context.Context, targetVersion string) (*atlasexec.MigrateDown, error) {
	client := dbm.atlasClient
	config := dbm.app.GetConfig()
	dbConn := config.DBConfig.Connection
	devDbConn := config.DevDBConfig.Connection

	params := &atlasexec.MigrateDownParams{
		URL:    dbConn.AtlasDSN(),
		DevURL: devDbConn.AtlasDSN(),
	}

	if targetVersion != "" {
		params.ToVersion = targetVersion
	}

	return client.MigrateDown(ctx, params)
}

func (dbm *AtlasDBMigrator) RehashMigration(ctx context.Context) error {
	client := dbm.atlasClient

	return client.MigrateHash(ctx, &atlasexec.MigrateHashParams{})
}

func (dbm *AtlasDBMigrator) PendingMigrationCheck(ctx context.Context) error {
	migrateStatus, err := dbm.GetMigrationStatus(ctx)
	if err != nil {
		return err
	}

	if len(migrateStatus.Pending) > 0 {
		return PendingMigrationErrors(migrateStatus.Pending)
	}

	return nil
}

func (dbm *AtlasDBMigrator) Close() error {
	return dbm.migrationDir.Close()
}
