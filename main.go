package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go conference" // variable - can change its value
	const conferenceTickets = 50      // constant - cannot change its value
	var remainingTickets uint8 = 50

	fmt.Printf("conferenceTickets is %T. remainingTickets is %T. conferenceName is %T.\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	bookings := []string{}
	var firstName string
	var lastName string
	var email string
	var userTickets uint8

	for true {
		fmt.Println("Please enter first name")
		fmt.Scan(&firstName) // Pass by reference. Will ask and read for user input, then store that value in the memory address

		fmt.Println("Please enter last name")
		fmt.Scan(&lastName) // Pass by reference. Will ask and read for user input, then store that value in the memory address

		fmt.Println("Please enter email")
		fmt.Scan(&email) // Pass by reference. Will ask and read for user input, then store that value in the memory address

		fmt.Println("Please enter number of tickets")
		fmt.Scan(&userTickets) // Pass by reference. Will ask and read for user input, then store that value in the memory address

		isValidName := len(firstName) >= 2 && len(lastName) >= 2 && (firstName == "ra" || firstName == "qq")
		isValidEmail := strings.Contains(email, "@")
		isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidUserTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("User %v %v booked %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings { // _ is a blank identified
				var names = strings.Fields(booking) // declare that it is a variable
				firstName = names[0]
				firstNames = append(firstNames, firstName)
			}
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
