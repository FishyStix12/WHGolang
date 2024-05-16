///////////////////////////////////////////////////////////////////////////////////////////////////////////
// Author: Nicholas Fisher
// Date: May 16 2024
// Description of the Script
// This Go script is designed to quickly scan for open TCP ports on the host `scanme.nmap.org`. It targets 
// ports from 1 to 1024, using goroutines to perform the scan concurrently, which significantly speeds up 
// the process. For each port, the script tries to establish a TCP connection. If the connection succeeds,
// the port is marked as open, and its number is printed to the console. Each connection attempt runs in 
// its own goroutine, allowing the script to efficiently complete the port scan. To test the speed of the 
// script, you can run it and measure the time it takes with the command `time ./toofast.go`.
///////////////////////////////////////////////////////////////////////////////////////////////////////////
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
