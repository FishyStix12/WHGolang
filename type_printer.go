///////////////////////////////////////////////////////////////////////////////////////////////////////////
// Author: Nicholas Fisher
// Date: March 18th 2024
// Description of the Script
//  The code in `type_printer.go` is a simple Go script that defines a function called `foo` which takes 
// an interface argument and prints a message indicating the type of the argument. It utilizes a type switch
// to handle different types of input: integers, strings, and booleans. The `main` function demonstrates the
// usage of the `foo` function with various types of inputs and prints out the corresponding type messages.
// This script serves as a basic illustration of type assertion and type switching in Go.
///////////////////////////////////////////////////////////////////////////////////////////////////////////
package main // Declares the package name. A Go program starts running in package main.
package main // Package main declares the package name. A Go program starts running in package main.

import (
	"fmt" // Imports the fmt package, which implements formatted I/O.
)

// foo prints the type of the given interface value.
func foo(i interface{}) {
	switch v := i.(type) { // Type switch: checks the type of the interface value i
	case int: // If i is of type int
		fmt.Printf("I'm an integer: %d\n", v) // Prints the value of i as an integer
	case string: // If i is of type string
		fmt.Printf("I'm a string: %s\n", v) // Prints the value of i as a string
	case bool: // If i is of type bool
		fmt.Printf("I'm a bool: %t\n", v) // Prints the value of i as a boolean
	default: // If i is of any other type
		fmt.Printf("Unknown type!\n") // Prints that the type is unknown
	}
}

func main() {
	// Sample usage of the foo function
	foo(42)             // Output: I'm an integer: 42
	foo("hello")        // Output: I'm a string: hello
	foo(3.14)           // Output: Unknown type!
	foo(true)           // Output: I'm a bool: true
	foo([]int{1, 2, 3}) // Output: Unknown type!
}
