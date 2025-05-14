package db

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/leeliwei930/walletassignment/constant"
	"github.com/leeliwei930/walletassignment/internal/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// dbMigrateRollbackCmd which allow to interact with atlas CLI to perform a down migrations
var dbMigrateRollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback database migrations",
	Long:  `This command reverts applied database migrations using Atlas CLI. It can rollback to a specific version and requires explicit confirmation when executed in production environments.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		targetVersion, _ := cmd.Flags().GetString("target-version")

		app, err := app.InitializeFromEnv()

		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		defer app.Close()

		dbMigrator := app.GetDBMigrator()
		defer dbMigrator.Close()

		environment := viper.GetString(constant.APP_ENV)
		slog.Info("Execute down migration on %s environment to version: %s", environment, targetVersion)

		rollbackMigrationResult, err := dbMigrator.RollbackMigration(ctx, targetVersion)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}

		messages := ""
		if len(rollbackMigrationResult.Reverted) == 0 {
			slog.Info("Atlas: No migrations can be reverted")
			return
		} else {
			messages += fmt.Sprintf("Atlas: Sucessfully revert migration to: %s", rollbackMigrationResult.Target)
			messages += "===\n"
		}

		if len(rollbackMigrationResult.Error) > 0 {
			messages += fmt.Sprintf("Atlas: Migration Down Errors - %s", rollbackMigrationResult.Error)
		}

		slog.Info(messages)
	},
}

func init() {
	dbMigrateRollbackCmd.Flags().String("target-version", "", "Revert the migrations to specified version")
	DBMigratorCmd.AddCommand(dbMigrateRollbackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbmigrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbmigrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
