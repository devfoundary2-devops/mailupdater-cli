package cmd

import (
	"fmt"
	"mailupdater/db"
	"mailupdater/utils"
	"os"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks if an email address is in the database",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check if argument length less that or greater than 2
		if len(args) != 2 {
			fmt.Print("Only two email addresses are required")
			os.Exit(1)
		}

		// validate email addresses
		oldMail, newMail := args[0], args[1]
		validOldMail := utils.ValidateMail(oldMail)
		validNewMail := utils.ValidateMail(newMail)
		if !validOldMail {
			fmt.Printf("Old email address provided is invalid: %s", oldMail)
			os.Exit(1)
		}
		if !validNewMail {
			fmt.Printf("New email address provided is invalid: %s", newMail)
			os.Exit(1)
		}

		// Connect to database
		database, err := db.NewDB()
		if err != nil {
			fmt.Printf("Error connecting to database: %v\n", err)
			os.Exit(1)
		}
		defer database.Close()

		// Check if old email exists in the database
		oldUser, err := database.CheckEmailsExists(oldMail)
		if err != nil {
			fmt.Printf("Error checking old email: %v\n", oldMail)
			os.Exit(1)
		}

		if oldUser == nil {
			fmt.Printf("Error: Old email '%s' not found in database\n", oldMail)
		}

		// Check if new email exists in the database
		newUser, err := database.CheckEmailsExists(newMail)
		if err != nil {
			fmt.Printf("Error checking new email: %v\n", newMail)
			os.Exit(1)
		}

		if newUser != nil {
			fmt.Print("Error: Both email addresses found in database!\n")
			fmt.Printf(" Old email: %s (User: %s %s)\n", oldUser.Email, oldUser.FirstName, oldUser.LastName)
			fmt.Printf(" New email: %s (User: %s %s)\n", newUser.Email, newUser.FirstName, newUser.LastName)
			fmt.Println("Cannot proceed with update.")
			os.Exit(1)
		}

		fmt.Println("User validation passed")
		fmt.Println("Ready to proceed with update.")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
