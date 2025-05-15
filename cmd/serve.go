package cmd

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/leeliwei930/walletassignment/server"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Initiate the pawswinq application's backend server",
	Long:  `Launches the backend server for the pawswinq application. The server will begin on the port defined in the .env file.`,
	Run: func(cmd *cobra.Command, args []string) {

		server.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
