package main

import (
	"GOProject/helper"
	"fmt"
	"sync"
	"time"
	// "strconv"
	// "strings"
)

var conferenceName = "Go Conference" // Define a variable
const conferenceTickets = 50

var remainingTickets = 50

// var bookings [50] string // Arrays of 50 elements
// var bookings = make([]map[string]string, 0) // Create list of maps, with initial size 0
var bookings = make([]UserData, 0) // Create list of maps, with initial size 0

type UserData struct {
	firstName string 
	lastName string 
	email string
	numberOfTickets int
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// fmt.Printf("conferenceTickets type : %T\n\n", conferenceTickets)

	// Create a loop to always ask again
	// for remainingTickets > 0 && len(bookings)<50 {
	// for {
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

			wg.Add(1)
			go sendTickets(userTickets, firstName, lastName, email) // "go" is used to eecute the code in another thread to optimize the app

			// Get all first names
			firstNames := getFirstNames()
			fmt.Printf("All bookings : %v \n", firstNames)

			// var noTicketsRemaining := remainingTickets == 0
			if remainingTickets <= 0 {
				// end program
				fmt.Println("Conference is booked out")
				// break
			}
		} else {
			fmt.Println("Invalid data")
		}
	// }
	wg.Wait() // Wait for the end of all threads before ending the program

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
		// var names = strings.Fields(booking)
		// firstNames = append(firstNames, names[0])
		// firstNames = append(firstNames, booking["firstName"]) // If we use map
		firstNames = append(firstNames, booking.firstName)
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

	// Create a map for a user
	// var userData = make(map[string]string) // We cannot mix types in the map => key must be string and value must be string
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatInt(int64(userTickets), 10) // 10 is for base 10

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Your first name : %v, your lastname : %v, your email : %v\n", firstName, lastName, email)
	fmt.Printf("It remains only %v tickets for %v\n", remainingTickets, conferenceName)
}

func sendTickets(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second) // Stop execution for 10 secondes
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("######")
	fmt.Printf("Sending ticket:\n %v \nto email adress %v\n", ticket, email)
	fmt.Println("######")
	wg.Done()
}