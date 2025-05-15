package db

import (
	"github.com/spf13/cobra"
)

var BuildVersion = "main"

var DBMigratorCmd = &cobra.Command{
	Use:   "db",
	Short: "Database migration commands",
	Long:  `This command group allows you to interact with the Atlas migration CLI to apply or revert migrations.`,
}
