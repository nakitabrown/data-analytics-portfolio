package main

import (
	"fmt" //print
	"strings"
)

// PACKAGE LEVEL VARIABLES AVAILABLE TO ALL FUNCS
// conferenceName := "Go Conference" | doesn't work with constants nor at package level variables
// Go can determine the data type without declaring it
const conferenceTickets int = 50

var conferenceName string = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

// MAIN FUNCTION
func main() {

	//GREETING FUNCTION
	greetUsers()

	for {
		//USER INPUT FUNCTION
		firstName, lastName, email, userTickets := getUserInput()

		//INPUT VALIDATION FUNCTION
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			//BOOK TICKETS FUNCTION
			bookTicket(userTickets, firstName, lastName, email)

			//FIRST NAMES FUNCTION
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break //break for loop
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}

			if !isValidEmail {
				fmt.Println("Email address you entered does not contain @ sign.")
			}

			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid.")
			}

		}

	}

}

// FUNCTIONS
func greetUsers() {
	//variable can be the same name as parameter or different
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

// []STRING IS OUTPUT
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames //PRINT FIRST NAMES - OUTPUT
}

// INPUT VALIDATION
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber //GO CAN RETURN MULTIPLE VALUES
}

// USER INPUT
// CREATE VARIABLES LOCAL AS POSSIBLE
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//USER INPUT
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	//SLICE
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for the %v.\n", remainingTickets, conferenceName)
	return

}
