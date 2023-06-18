package main

import (
	"fmt"
	"strconv"
)

// package level variables
var conferenceName = "Go conference" // variable - can change its value
const conferenceTickets = 50         // constant - cannot change its value
var remainingTickets uint8 = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidUserTickets := ValidateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidUserTickets {
			remainingTickets, bookings = bookTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("The first names of our bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("All tickets are sold out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("User's first name or last name entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email entered is invalid")
			}
			if !isValidUserTickets {
				fmt.Println("Number of tickets is invalid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { // _ is a blank identifier
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint8) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint8

	fmt.Println("Please enter first name")
	fmt.Scan(&firstName) // Pass by reference. Will ask and read for user input, then store that value in the memory address

	fmt.Println("Please enter last name")
	fmt.Scan(&lastName) // Pass by reference. Will ask and read for user input, then store that value in the memory address

	fmt.Println("Please enter email")
	fmt.Scan(&email) // Pass by reference. Will ask and read for user input, then store that value in the memory address

	fmt.Println("Please enter number of tickets")
	fmt.Scan(&userTickets) // Pass by reference. Will ask and read for user input, then store that value in the memory address

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint8, firstName string, lastName string, email string) (uint8, []map[string]string) {
	remainingTickets = remainingTickets - userTickets

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v", bookings)

	fmt.Printf("User %v %v booked %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return remainingTickets, bookings
}
