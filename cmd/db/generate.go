package db

import (
	"context"
	"os"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/joho/godotenv/autoload"
	"github.com/leeliwei930/walletassignment/ent/migrate"
	"github.com/leeliwei930/walletassignment/internal/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// makeMigrateCmd represents the makeMigrate command
var makeMigrateCmd = &cobra.Command{
	Use:   "generate:migration <name>",
	Short: "Formulate a migration file from a specified database and determine the target state using ent schemas",
	Long:  `This command is responsible for creating SQL migration files from a specific database, and determining the target state using the present ent schemas`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runMakeMigrateCmd(cmd, args)
	},
}

func initConfig() {
	viper.AutomaticEnv()
}

func init() {

	cobra.OnInitialize(initConfig)

	makeMigrateCmd.PersistentFlags().StringP("mode", "m", "inspect", "The migration create mode default is inspect")
	DBMigratorCmd.AddCommand(makeMigrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// makeMigrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// makeMigrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func runMakeMigrateCmd(cmd *cobra.Command, args []string) {

	migrationName := &args[0]

	app, err := app.InitializeFromEnv()
	if err != nil {
		panic(err)
	}

	// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
	dbConfig := app.GetConfig().DevDBConfig
	log := app.GetLog()
	ctx := context.Background()
	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir("database/migrations")
	if err != nil {
		log.Error("failed creating atlas migration directory", zap.Error(err))
		os.Exit(1)
	}

	devDBConfig := app.GetConfig().DevDBConfig
	dialect := devDBConfig.Connection.EntDialect()

	var migrationMode schema.Mode
	migrationMode = schema.ModeInspect
	migrationModeStr, _ := cmd.Flags().GetString("mode")
	if migrationModeStr == "replay" {
		migrationMode = schema.ModeReplay
	}
	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                     // provide migration directory
		schema.WithMigrationMode(migrationMode), // provide migration mode
		schema.WithDialect(dialect),             // Ent dialect to use
		schema.WithIndent("    "),
		schema.WithFormatter(atlas.DefaultFormatter),
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
		schema.WithSkipChanges(
			schema.DropTable,
		),
	}

	if len(*migrationName) == 0 {
		log.Error("migration name is required. Use: 'go run -mod=mod cmd/migrate.go --name <name>'")
		os.Exit(1)
	}

	err = migrate.NamedDiff(ctx, dbConfig.Connection.AtlasDSN(), *migrationName, opts...)
	if err != nil {
		log.Error("failed generating migration file", zap.Error(err))
		os.Exit(1)
	}
}
