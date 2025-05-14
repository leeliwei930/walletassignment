package db

import (
	"context"
	"log/slog"
	"os"

	"github.com/leeliwei930/walletassignment/internal/app"
	"github.com/spf13/cobra"
)

var rehashCommand = &cobra.Command{
	Use:   "rehash",
	Short: "Rehash the atlas.sum file based on current migration files",
	Long: `This command proxies to Atlas CLI's 'migrate rehash' command, which recalculates the checksum 
of migration files and updates the atlas.sum file. 
Use this after manually modifying migration files or when resolving checksum verification failures.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		app, err := app.InitializeFromEnv()
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		}
		defer app.Close()

		dbMigrator := app.GetDBMigrator()
		defer dbMigrator.Close()

		err = dbMigrator.RehashMigration(ctx)
		if err != nil {
			slog.Error(err.Error())
			os.Exit(1)
		} else {
			slog.Info("Atlas: Rehashing successful")
		}

	},
}

func init() {
	DBMigratorCmd.AddCommand(rehashCommand)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbmigrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbmigrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
