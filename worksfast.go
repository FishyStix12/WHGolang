///////////////////////////////////////////////////////////////////////////////////////////////////////////
// Author: Nicholas Fisher
// Date: May 16 2024
// Description of the Script
// The provided Go script is a concurrent port scanner that checks the availability of ports from 1 to 
// 1024 on the host "scanme.nmap.org". It utilizes goroutines to handle each port scan concurrently,
// significantly speeding up the process. A `sync.WaitGroup` is used to ensure the main function waits for
// all goroutines to complete before exiting. For each port, the script attempts to establish a TCP 
// connection; if successful, it prints the port number and closes the connection. Errors, such as 
// connection refusals, are handled gracefully by simply returning from the goroutine. This approach 
// demonstrates efficient use of Go's concurrency features to perform network operations.
///////////////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	// Create a wait group to synchronize the completion of goroutines
	var wg sync.WaitGroup

	// Loop through port numbers from 1 to 1024
	for i := 1; i <= 1024; i++ {
		// Increment the wait group counter for each port number
		wg.Add(1)

		// Start a new goroutine for each port number
		go func(j int) {
			// Decrement the wait group counter when the goroutine completes
			defer wg.Done()

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
		}(i) // Pass the current value of i to the goroutine
	}

	// Wait for all goroutines to complete
	wg.Wait()
}
