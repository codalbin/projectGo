package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference" // Define a variable
	const conferenceTickets = 50
	var remainingTickets = 50

	// fmt.Printf("conferenceTickets type : %T\n\n", conferenceTickets)

	// Change the %v if it is something else than text (number for example)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("Get your ticket, there are only", conferenceTickets, "place and there are only", remainingTickets, "still available")

	// var bookings [50] string // Arrays of 50 elements
	var bookings []string
	var firstName string
	var lastName string
	var email string
	var userTickets int

	// Create a loop to always ask again
	// for remainingTickets > 0 && len(bookings)<50 {
	for {
		// Ask user for their name
		fmt.Printf("Enter your first name\n")
		fmt.Scan(&firstName) // Use pointer

		fmt.Printf("Enter your last name\n")
		fmt.Scan(&lastName) // Use pointer

		fmt.Printf("Enter your email\n")
		fmt.Scan(&email) // Use pointer

		fmt.Printf("How many tickets do you want ?\n")
		fmt.Scan(&userTickets) // Use pointer

		var isValidName = len(firstName) >= 2 && len(lastName) >= 2
		var isValidEmail = strings.Contains(email, "@")
		isValidUserTickets := userTickets > 0 && userTickets <= remainingTickets

		// if userTickets > remainingTickets {
		// 	fmt.Printf("We only have only %v remaining, you can't book tickets\n", remainingTickets)
		// 	continue // Skip the end and move to the next iteration
		// }
		// else if {}
		// else {}

		if isValidName && isValidEmail && isValidUserTickets {
			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Your first name : %v, your lastname : %v, your email : %v\n", firstName, lastName, email)
			fmt.Printf("It remains only %v tickets\n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings { // The _ corresponds to the index
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("All bookings : %v \n", firstNames)

			// var noTicketsRemaining := remainingTickets == 0
			if remainingTickets <= 0 {
				// end program
				fmt.Println("Conference is booked out")
				break
			}
		} else {
			fmt.Println("Invalid data")
		}
	}

	// city := "London"

	// switch city {
	// 	case "New York" :
	// 		// code if city is NY
	// 	case "London" :
	// 		//
	// 	case "Paris", "Berlin" :
	// 		//
	// 	default :
	// 		fmt.Print("No valid city selected")

	// }
}
