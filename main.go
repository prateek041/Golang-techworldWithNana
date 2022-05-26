// All our code must belong to a package.
package main // standard name for the main package is "main".

// all the packages that are needed for the program are needed to be explicitly imported.
import "fmt"

// Go needs an entry point to start executing the program. It looks for a main function
// to begin with.
func main() { // There can only be one main function per application.
	fmt.Println("Hello world") // "Print" is a function in the fmt package.

}
