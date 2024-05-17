///////////////////////////////////////////////////////////////////////////////////////////////////////////
// Author: Nicholas Fisher
// Date: May 16 2024
// Description of the Script
// The script is a Go program that performs a basic port scan on the domain `scanme.nmap.org` to identify
// which ports (from 1 to 1024) are open. It utilizes goroutines for concurrent execution, significantly
// speeding up the scanning process. The `worker` function tries to establish a TCP connection to each port,
// sending the port number to a results channel if the connection is successful, otherwise sending a zero.
// The `main` function initializes channels for port numbers and results, spawns multiple worker goroutines,
// and feeds port numbers into the ports channel. It then collects results, filters out closed ports, sorts 
// the list of open ports, and prints each open port. This parallel approach enables efficient scanning and
// demonstrates the power of concurrent programming in Go.
///////////////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"fmt"  // Importing the fmt package for formatted I/O
	"net"  // Importing the net package for network I/O
	"sort" // Importing the sort package for sorting slices
)

// The worker function scans ports by attempting to establish a TCP connection.
// It reads port numbers from the ports channel and sends results to the results channel.
func worker(ports, results chan int) {
	// Iterating over each port received from the ports channel
	for p := range ports {
		// Formatting the address string with the current port number
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		// Attempting to establish a TCP connection to the address
		conn, err := net.Dial("tcp", address)
		// If there is an error (connection failed), send 0 to results and continue
		if err != nil {
			results <- 0
			continue
		}
		// If connection is successful, close it and send the port number to results
		conn.Close()
		results <- p
	}
}

func main() {
	// Creating a buffered channel to hold port numbers (up to 200)
	ports := make(chan int, 300)
	// Creating a channel to hold scan results
	results := make(chan int)
	// Slice to store open port numbers
	var openports []int

	// Starting worker goroutines (one for each slot in the buffered ports channel)
	for i := 0; i <= cap(ports); i++ {
		go worker(ports, results)
	}

	// Anonymous goroutine to send port numbers (1 to 1024) to the ports channel
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	// Receiving results from the results channel and storing open ports in openports slice
	for i := 0; i < 1024; i++ {
		port := <-results
		// If the port is open (non-zero), append it to the openports slice
		if port != 0 {
			openports = append(openports, port)
		}
	}

	// Closing the channels (optional here as the program is about to exit)
	close(ports)
	close(results)

	// Sorting the open ports in ascending order
	sort.Ints(openports)
	// Printing each open port
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
