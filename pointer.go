///////////////////////////////////////////////////////////////////////////////////////////////////////////
// Author: Nicholas Fisher
// Date: March 18th 2024
// Description of the Script
//  The provided Go script demonstrates the usage of pointers in Go programming language. It initializes 
// an integer variable `count` with the value 42, creates a pointer `ptr` to this variable, and then prints
// the value of `count` via the pointer. Afterward, it modifies the value of `count` through the pointer,
// assigning it the value 100, and prints the updated value of `count`. This script illustrates how 
// pointers can be used to indirectly access and modify variables in Go.
///////////////////////////////////////////////////////////////////////////////////////////////////////////
package main // Declares the package name. A Go program starts running in package main.

import "fmt" // Imports the fmt package, which implements formatted I/O.

func main() { // Declares the main function, which serves as the entry point of the program.
	var count = int(42) // Declares a variable named "count" and initializes it with the value 42.
	ptr := &count // Declares a pointer variable named "ptr" and assigns it the memory address of the "count" variable.
	fmt.Println(*ptr) // Prints the value pointed to by "ptr", which is the value of "count".
	*ptr = 100 // Modifies the value pointed to by "ptr" to 100, effectively changing the value of "count".
	fmt.Println(count) // Prints the new value of the "count" variable, which is now 100.
}
