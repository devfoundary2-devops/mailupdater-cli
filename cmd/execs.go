package cmd

import (
	"errors"
	"fmt"
	"mailupdater/db"
	"mailupdater/utils"
)

func CheckMails(args []string) (bool, error) {
	oldMail, newMail := args[0], args[1]
	validOldMail := utils.ValidateMail(oldMail)
	validNewMail := utils.ValidateMail(newMail)
	if !validOldMail {
		return false, errors.New("old email is invalid")
	}
	if !validNewMail {
		return false, errors.New("new email is invalid")
	}

	// Connect to database
	database, err := db.NewDB()
	if err != nil {
		return false, errors.New("error connecting to database")
	}
	defer database.Close()

	// Check if old email exists in the database
	oldUser, err := database.CheckEmailsExists(oldMail)
	if err != nil {
		return false, errors.New("error checking old email")
	}

	if oldUser == nil {
		return false, errors.New("error old email not found in database")
	}

	// Check if new email exists in the database
	newUser, err := database.CheckEmailsExists(newMail)
	if err != nil {
		return false, errors.New("error checking new email")
	}

	if newUser != nil {
		fmt.Print("Error: Both email addresses found in database!\n")
		fmt.Printf(" Old email: %s (User: %s %s)\n", oldUser.Email, oldUser.FirstName, oldUser.LastName)
		fmt.Printf(" New email: %s (User: %s %s)\n", newUser.Email, newUser.FirstName, newUser.LastName)
		fmt.Println("Cannot proceed with update.")
		return false, errors.New("both emails found in database")
	}

	fmt.Println("User validation passed")
	return true, nil
}
