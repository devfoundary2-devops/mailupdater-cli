package cmd

import (
	"fmt"
	"os"
	"regexp"

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
		validOldMail := validateMail(oldMail)
		validNewMail := validateMail(newMail)
		if !validOldMail {
			fmt.Printf("Old email address provided is invalid: %s", oldMail)
			os.Exit(1)
		}
		if !validNewMail {
			fmt.Printf("New email address provided is invalid: %s", newMail)
			os.Exit(1)
		}

		// check if old and new email addresseses are in the db
		// connect to the database, get a query cursor

		// respond with validation message
	},
}

func validateMail(s string) bool {
	emailRegex := `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(s)
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
