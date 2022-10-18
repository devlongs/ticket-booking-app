package main

import (
	"fmt"
)

func main (){
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	// ask user for their first name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	// ask user for their last name
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	// ask user for their email
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	// ask user for number of tickets
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
}