package main

import (
	"fmt"
	"strings"
)

func main (){
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
    bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isVaildEmail := strings.Contains(email, "@")
		isValidTIcketNumber := userTickets > 0 &&  userTickets <= remainingTickets


		// book ticket logic
		if isValidName && isVaildEmail && isValidTIcketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Comeback next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isVaildEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTIcketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
		
	}

}