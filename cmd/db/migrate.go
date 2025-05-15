package db

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/leeliwei930/walletassignment/constant"
	"github.com/leeliwei930/walletassignment/internal/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// dbmigrateCmd which allow to interact with atlas CLI to apply migrations
var dbmigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Apply pending database migrations",
	Long:  `This command acts as a proxy to the Atlas migration CLI, applying pending migrations to the database. It supports dry-run mode and targeting specific versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		isDryRun, _ := cmd.Flags().GetBool("dry-run")
		targetVersion, _ := cmd.Flags().GetString("target-version")
		envFile, _ := cmd.Flags().GetString("env-file")

		err := godotenv.Overload(envFile)
		if err != nil {
			slog.Error(err.Error())
		}

		app, err := app.InitializeFromEnv()
		if err != nil {
			slog.Error(err.Error())
		}

		dbMigrator := app.GetDBMigrator()
		defer dbMigrator.Close()

		environment := viper.GetString(constant.APP_ENV)
		if isDryRun {
			slog.Info("Execute migration in dry-run mode on", "environment", environment)
		} else {
			slog.Info("Execute migration", "environment", environment)
		}

		applyMigrationResult, err := dbMigrator.ApplyMigration(ctx, targetVersion, isDryRun)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}

		messages := ""
		if len(applyMigrationResult.Applied) == 0 {
			slog.Info("Atlas: No new migrations to apply")
			os.Exit(0)
		} else {
			messages += "Atlas: Applied Migrations: \n"
			for idx, migration := range applyMigrationResult.Applied {
				messages += fmt.Sprintf("[%d] Version: %s, Name: %s \n", idx+1, migration.Version, migration.Name)
			}
			messages += "===\n"
		}

		if len(applyMigrationResult.Error) > 0 {
			messages += "Atlas: Migration Errors: \n"
			messages += applyMigrationResult.Error + "\n"
			messages += "===\n"
		}

		messages += applyMigrationResult.Summary("    ")
		slog.Info(messages)
	},
}

func init() {
	dbmigrateCmd.Flags().Bool("dry-run", false, "Perform migration in dry-run mode, no changes will be applied to the database")
	dbmigrateCmd.Flags().String("target-version", "", "Apply migrations up to the specified version")
	dbmigrateCmd.Flags().String("env-file", ".env", "Environment file to use")
	DBMigratorCmd.AddCommand(dbmigrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbmigrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbmigrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
