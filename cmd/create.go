/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"

	"github.com/leeliwei930/walletassignment/internal/app"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new user",
	Long:  `Create a new user with the given phone number`,
	Run: func(cmd *cobra.Command, args []string) {
		app, err := app.InitializeFromEnv()
		if err != nil {
			log.Fatalf("Failed to initialize application: %v", err)
		}
		defer app.Close()

		userSvc := app.GetUserService()
		log := app.GetLog()

		phoneNumber, _ := cmd.Flags().GetString("phoneNumber")
		firstName, _ := cmd.Flags().GetString("fname")
		lastName, _ := cmd.Flags().GetString("lname")

		if len(phoneNumber) == 0 {
			log.Error("Phone number is required")
			return
		}
		if len(firstName) == 0 {
			log.Error("First name is required")
			return
		}
		if len(lastName) == 0 {
			log.Error("Last name is required")
			return
		}

		user, err := userSvc.SetupUser(context.Background(), phoneNumber, firstName, lastName)
		if err != nil {
			log.Error("Failed to create user", zap.Error(err))
			return
		}

		log.Info("User created successfully", zap.Any("user", user))
	},
}

func init() {
	createCmd.Flags().StringP("phoneNumber", "p", "", "Phone number of the user")
	createCmd.Flags().StringP("fname", "f", "", "First name of the user")
	createCmd.Flags().StringP("lname", "l", "", "Last name of the user")

	userCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
