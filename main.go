// All our code must belong to a package.
package main // standard name for the main package is "main".

// all the packages that are needed for the program are needed to be explicitly imported.
import "fmt"

// Go needs an entry point to start executing the program. It looks for a main function
// to begin with.
func main() { // only one main fucntion per program.
	// fmt.Println("Hello world") // "Print": function in the fmt package.

	// var conferenceName = "Go Conference"                               // Variable.
	conferenceName := "Go conference"                                     // shorthand for variale declaration/assignment. has its restrictions.
	const conferenceTickets uint = 50                                     // Constant.
	var remainingTickets uint = 50                                        // Go can assign the type, depending upon the data assigned.
	fmt.Printf("welcome to the %v booking application\n", conferenceName) // %v is the placeholder
	fmt.Println("Get your tickets here")
	fmt.Printf("%v tickets remaining from %v tickets\n", remainingTickets, conferenceTickets)

	// var bookings = [50] string{} // array : bookings, size : 50, data-type: string, currently empty.
	var bookings [50]string // shorthand.
	var firstName string
	var lastName string
	var email string
	var userTicket uint

	// ask for the user input
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName) // scan user input and assign to userName variable.

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want")
	fmt.Scan(&userTicket)

	// stroing the values in the database.
	bookings[0] = firstName + " " + lastName
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("The array type: %T\n", bookings)
	fmt.Printf("The array length: %v\n", len(bookings))

	// booking ticket logic.
	remainingTickets -= userTicket // shorthand for remainingTickets = remainingTickets - userTicket.

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining\n", remainingTickets)
}
