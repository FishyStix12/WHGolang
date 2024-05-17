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
