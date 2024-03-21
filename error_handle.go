///////////////////////////////////////////////////////////////////////////////////////////////////////////
// Author: Nicholas Fisher
// Date: March 18th 2024
// Description of the Script
// The code defines a custom error type MyError, implements the error interface for it, and creates a
// function foo() that returns an error of this custom type. In the main() function, it calls foo() and 
// checks if an error is returned. If an error occurs, it prints out the error message.
///////////////////////////////////////////////////////////////////////////////////////////////////////////
package main // Declares the package name. A Go program starts running in package main.

import (
	"fmt"
)

// error is an interface type that defines the method Error().
// Any type that implements this method can be treated as an error.
type MyError string

// Error method returns the string representation of the error.
func (e MyError) Error() string {
	return string(e)
}

// foo function returns an error of type MyError.
func foo() error {
	return MyError("Some Error Occurred")
}

func main() {
	// Call foo() function which may return an error.
	if err := foo(); err != nil {
		// Handle the error by printing it.
		fmt.Println("Error:", err)
		// In a real-world scenario, you might perform additional error handling here.
	}
}
