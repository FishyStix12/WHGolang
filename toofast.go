package main

import (
	"fmt"
	"net"
)

func main() {
	// Loop through port numbers from 1 to 1024
	for i := 1; i <= 1024; i++ {
		// Start a new goroutine for each port number
		go func(j int) {
			// Create the address string for the port scan
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			// Try to establish a TCP connection to the address
			conn, err := net.Dial("tcp", address)
			if err != nil {
				// If there's an error (e.g., connection refused), simply return
				return
			}
			// If the connection is successful, close it immediately
			conn.Close()
			// Print the port number if it is open
			fmt.Printf("%d open\n", j)
		}(i)
	}
}
