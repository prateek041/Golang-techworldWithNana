// All our code must belong to a package.
package main // standard name for the main package is "main".

// all the packages that are needed for the program are needed to be explicitly imported.
import "fmt"

// Go needs an entry point to start executing the program. It looks for a main function
// to begin with.
func main() { // There can only be one main function per application.
	// fmt.Println("Hello world") // "Print" is a function in the fmt package.

	// var conferenceName = "Go Conference"                               // Variable.
	conferenceName := "Go conference"                                     // shorhand for variale declaration/assignment. has its restrictions.
	const conferenceTickets uint = 50                                     // Constant.
	var remainingTickets uint = 50                                        // Go can assign the type, depending upon the data assigned.
	fmt.Printf("welcome to the %v booking application\n", conferenceName) // %v is the placeholder
	fmt.Println("Get your tickets here")
	fmt.Printf("%v tickets remaining from %v tickets\n", remainingTickets, conferenceTickets)

	var userName string
	var userTicket int
	// ask for the user input

	userName = "prateek"
	userTicket = 5
	fmt.Printf("The user %v booked %v tickets\n", userName, userTicket) // we can use %T to print the type
	// of the given variable.
	fmt.Printf("conferenceTicket is of type %T, conferenceName is of type %T\n", conferenceTickets, conferenceName)
}
