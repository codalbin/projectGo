package main

import (
	"fmt"
	"strings"
	"GOProject/helper"
)

var conferenceName = "Go Conference" // Define a variable
const conferenceTickets = 50

var remainingTickets = 50

// var bookings [50] string // Arrays of 50 elements
var bookings = []string{}

func main() {

	greetUsers()

	// fmt.Printf("conferenceTickets type : %T\n\n", conferenceTickets)

	// Create a loop to always ask again
	// for remainingTickets > 0 && len(bookings)<50 {
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// if userTickets > remainingTickets {
		// 	fmt.Printf("We only have only %v remaining, you can't book tickets\n", remainingTickets)
		// 	continue // Skip the end and move to the next iteration
		// }
		// else if {}
		// else {}

		if isValidName && isValidEmail && isValidUserTickets {

			bookTicket(userTickets, firstName, lastName, email)
			// Get all first names
			firstNames := getFirstNames()
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

func greetUsers() {
	// Change the %v if it is something else than text (number for example)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("Get your ticket, there are only", conferenceTickets, "place and there are only", remainingTickets, "still available")

}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { // The _ corresponds to the index
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	// Ask user for their name
	fmt.Printf("Enter your first name\n")
	fmt.Scan(&firstName) // Use pointer

	fmt.Printf("Enter your last name\n")
	fmt.Scan(&lastName) // Use pointer

	fmt.Printf("Enter your email\n")
	fmt.Scan(&email) // Use pointer

	fmt.Printf("How many tickets do you want ?\n")
	fmt.Scan(&userTickets) // Use pointer

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Your first name : %v, your lastname : %v, your email : %v\n", firstName, lastName, email)
	fmt.Printf("It remains only %v tickets for %v\n", remainingTickets, conferenceName)
}
