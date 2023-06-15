package main

import "fmt"

func main() {

	var conferenceName = "Go conference" // variable - can change its value
	const conferenceTickets = 50         // constant - cannot change its value
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")

}
