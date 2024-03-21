///////////////////////////////////////////////////////////////////////////////////////////////////////////
// Author: Nicholas Fisher
// Date: March 18th 2024
// Description of the Script
//  The script demonstrates the use of goroutines and channels in Go. It defines a function f() 
// that prints a message, and a function strlen() that calculates the length of a string and sends it to a 
// channel. In the main() function, f() is executed concurrently in a goroutine, followed by a 1-second sleep.
// Then, the length of the string "Salutations" is calculated concurrently using strlen() and sent to a channel.
// The main function receives two values from the channel (the length of "Salutations" and 0), prints them,
// and prints their sum.
///////////////////////////////////////////////////////////////////////////////////////////////////////////
package main // Declares the package name. A Go program starts running in package main.
package main // Declares the package name

import (
	"fmt"      // Imports the fmt package for formatted I/O
	"time"     // Imports the time package for time-related functions
)

func f() { // Defines a function named f which doesn't take any arguments and doesn't return anything
	fmt.Println("f function") // Prints "f function"
}

func strlen(s string, c chan int) { // Defines a function named strlen which takes a string s and a channel c as arguments
	c <- len(s) // Sends the length of the string s to the channel c
}

func main() { // Defines the main function
	go f() // Launches the f function in a separate goroutine
	time.Sleep(1 * time.Second) // Pauses the main goroutine for 1 second

	fmt.Println("main function") // Prints "main function"

	c := make(chan int) // Creates a channel named c of type int
	go strlen("Salutations", c) // Launches the strlen function in a separate goroutine, passing "Salutations" and the channel c
	x, y := <-c, <-c // Receives values from the channel c into variables x and y
	fmt.Println(x, y, x+y) // Prints the values of x, y, and their sum
}
