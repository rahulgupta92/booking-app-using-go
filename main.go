package main

import "fmt"

func main() {

	conferenceName := "Go conference" // variable - can change its value
	const conferenceTickets = 50      // constant - cannot change its value
	var remainingTickets uint8 = 50

	fmt.Printf("conferenceTickets is %T. remainingTickets is %T. conferenceName is %T.\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

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

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("User %v %v booked %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v", remainingTickets, conferenceName)
}
