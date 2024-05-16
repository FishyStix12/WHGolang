///////////////////////////////////////////////////////////////////////////////////////////////////////////
// Author: Nicholas Fisher
// Date: May 16 2024
// Description of the Script
//  This Go script conducts a TCP port scan on the server "scanme.nmap.org" across ports 1 through 1024. 
// It utilizes concurrency to efficiently scan multiple ports simultaneously. For each port, it attempts 
// to establish a TCP connection. If the connection is successful, it indicates that the port is open; 
// otherwise, it assumes the port is either closed or filtered. The script prints the status of each port,
// either indicating it as open or specifying that it is closed or filtered.
///////////////////////////////////////////////////////////////////////////////////////////////////////////
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
