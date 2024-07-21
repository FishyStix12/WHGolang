///////////////////////////////////////////////////////////////////////////////////////////////////////////
// Author: Nicholas Fisher
// Date: March 18th 2024
// Description of the Script
//  The provided Go code demonstrates the usage of a for loop and range iteration in Go. It first prints
// numbers from 0 to 10 using a simple for loop, then it iterates over the elements of a slice (nums)
// using the range keyword, printing the index and value of each element.
///////////////////////////////////////////////////////////////////////////////////////////////////////////
package main // Declares the package name. A Go program starts running in package main.

import (
	"fmt" // Imports the fmt package, which implements formatted I/O.
)

func main() {
	// Loop to print numbers from 0 to 10
	for i := 0; i <= 10; i++ { // Initialize i to 0; continue looping until i is less than or equal to 10; increment i by 1 each iteration
		fmt.Println(i) // Print the value of i
	}

	nums := []int{2, 4, 6, 8} // Declare a slice of integers
	// Loop to iterate over elements of the nums slice
	for idx, val := range nums { // Iterate over the elements of the nums slice, where idx is the index and val is the value of each element
		fmt.Println(idx, val) // Print the index and value of the current element in the slice
	}
}
