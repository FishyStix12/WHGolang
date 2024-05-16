///////////////////////////////////////////////////////////////////////////////////////////////////////////
// Author: Nicholas Fisher
// Date: May 16 2024
// Description of the Script
// This Go script attempts to establish a TCP connection to the server "scanme.nmap.org" on port 80. If 
// the connection is successful, it prints "Connection successful" to the console.
///////////////////////////////////////////////////////////////////////////////////////////////////////////
package main // Declares that this file belongs to the main package, which can be executed as a standalone program.

import (
	"fmt" // Importing the fmt package for printing messages to the console.
	"net" // Importing the net package for network-related operations.
)

func main() { // Entry point of the program.
	_, err := net.Dial("tcp", "scanme.nmap.org:80") // Attempting to establish a TCP connection to "scanme.nmap.org" on port 80.
	if err == nil {                                 // Checking if there was no error during the connection attempt.
		fmt.Println("Connection successful") // If no error occurred, print "Connection successful".
	}
}
