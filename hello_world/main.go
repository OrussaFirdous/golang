// Go code is organised into packages just like libraries or modules in other languages
// Package main defines a standalone executable program,not a library
package main

// List of packages that the program imports
// Go allows to import packages that you excatly need.
// The program won't compile if there are missing imports or missing
// the unnecessary ones.
import (
	"fmt"
)

// Main function marks the beginning of execution of prgoram
func main() {
	fmt.Println("Hello World!")
}
