package cmd

import (
	"fmt"
	"mailupdater/db"
	"mailupdater/utils"
	"os"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [old-email] [new-email]",
	Short: "Updates a user's email address in the database",
	Long: `Updates a user's email address after validating that:
- The old email exists in the database
- The new email does not already exist
- Both emails are valid email addresses

Example:
	mailupdater update john.doe@example.com john.doe@hey.com`,
	Run: func(cmd *cobra.Command, args []string) {
		oldMail, newMail := args[0], args[1]

		if !utils.ValidateMail(oldMail) {
			fmt.Printf("Error: Old email address is invalid: %s\n", oldMail)
			os.Exit(1)
		}

		if !utils.ValidateMail(newMail) {
			fmt.Printf("Error: New email address is invalid: %s\n", newMail)
			os.Exit(1)
		}

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

		// Perform the mail update
		fmt.Printf("Updating email for %s %s...\n", oldUser.FirstName, oldUser.LastName)
		fmt.Printf(" From: %s\n", oldMail)
		fmt.Printf(" To: %s\n", newMail)

		err = database.UpdateEmail(oldMail, newMail)
		if err != nil {
			fmt.Printf("Error updating email: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Email updated successfully.")

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
