package main // Declares that this file belongs to the main package, which can be executed as a standalone program.

import (
	"fmt" // Importing the fmt package for printing messages to the console.
	"net" // Importing the net package for golang.
)

func main() {
	// For loop defines the ports we should do on each scan.
	for i := 1; i <= 1024; i++ {
		// uses the for function to do concurrent scans on scanme.nmap.org on ports 1-1024
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		// this line initiates the scan on the address variable
		conn, err := net.Dial("tcp", address)
		// If the err is not nil as it should be this is what the scan will do.
		if err != nil {
			//Port is closed or filtered, and this function will tell us which ports are closed or filtered
			fmt.Printf("Port %d is closed or filtered", i)
			// tells it to go to the next port to scan
			continue
		}
		// Closes the connection before going to the next port
		conn.Close()
		// Tell us the port is open.
		fmt.Printf("%d is open!", i)
	}
}
