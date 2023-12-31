package main

import "strings"

// variable or function shared across packages - global
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint8) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidUserTickets
}
