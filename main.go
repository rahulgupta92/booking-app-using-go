package main

import "fmt"

func main() {

	var conferenceName = "Go conference" // variable - can change its value
	const conferenceTickets = 50         // constant - cannot change its value
	var remainingTickets = 50

	fmt.Printf("conferenceTickets is %T. remainingTickets is %T. conferenceName is %T.\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int

	userName = "rahul"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
