// /////////////////////////////////////////////////////////////////////////////////////////////////////////
// Author: Nicholas Fisher
// Date: March 18th 2024
// Description of the Script
// This script demonstrates a simple concurrency example in Go. It defines a function strlen which calculates
// the length of a given string and sends it to a channel. Then, in the main function, it creates a channel,
// launches strlen in a separate goroutine, receives the length of the string from the channel, and prints
// the sum of the lengths of the string.\
// /////////////////////////////////////////////////////////////////////////////////////////////////////////
package main // Declares the package name. A Go program starts running in package main.
import (
	"fmt"
)

func strlen(s string, c chan int) { // Defines a function named strlen which takes a string s and a channel c as arguments
	c <- len(s) // Sends the length of the string s to the channel c
}

func main() { // Defines the main function
	c := make(chan int)         // Creates a channel named c of type int
	go strlen("Salutations", c) // Launches the strlen function in a separate goroutine, passing "Salutations" and the channel c
	go strlen("World", c)       // Launches the strlen function in a separate goroutine, passing "World" and the channel c
	x, y := <-c, <-c            // Receives values from the channel c into variables x and y
	fmt.Println(x, y, x+y)      // Prints the values of x, y, and their sum
}
