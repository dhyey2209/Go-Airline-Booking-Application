package helper

import (
	"strings"
)

// ValidateUserInput checks if the user's input meets the requirements for booking an airline seat.
func ValidateUserInput(firstName string, lastName string, email string, passportNumber string, userSeats uint, remainingSeats uint) (bool, bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidPassportNumber := validatePassportNumber(passportNumber)
	isValidSeatNumber := userSeats > 0 && userSeats <= remainingSeats
	return isValidName, isValidEmail, isValidPassportNumber, isValidSeatNumber
}

// validatePassportNumber checks if the passport number is of a valid format.
func validatePassportNumber(passportNumber string) bool {
	// The validity criteria for a passport number can vary by country and format.
	// Please implement the appropriate validation logic for your requirements.
	return len(passportNumber) > 0 // This is a placeholder for a real validation check.
}
