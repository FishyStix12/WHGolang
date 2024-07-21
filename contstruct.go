///////////////////////////////////////////////////////////////////////////////////////////////////////////
// Author: Nicholas Fisher
// Date: March 18th 2024
// Description of the Script
//  The script is a simple Go program that demonstrates the usage of conditionals (if statement) and switch 
// statements. It defines a variable x, assigns it the value "foo", and then checks whether x is equal to 
// "1" using an if statement. Depending on the value of x, it prints out a corresponding message. 
// Additionally, it employs a switch statement to determine the action based on the value of x, printing 
// different messages for different cases.
///////////////////////////////////////////////////////////////////////////////////////////////////////////
package main // Declares the package name. A Go program starts running in package main.

import "fmt" // Imports the fmt package, which implements formatted I/O.

func main() {
	x := "foo" // Define the variable x and assign it the value "foo"

	if x == "1" { // Check if the value of x is equal to the string "1"
		fmt.Println("X is equal to 1") // If the condition is true, print this message
	} else {
		fmt.Println("X is not equal to 1") // If the condition is false, print this message
	}

	// This is like a For Conditional involving more than two options
	switch x { // Start a switch statement based on the value of x
	case "foo": // If x is "foo"
		fmt.Println("Found foo") // Print this message
	case "bar": // If x is "bar"
		fmt.Println("Found bar") // Print this message
	default: // If x doesn't match any case
		fmt.Println("Default Case") // Print this message
	}
}
