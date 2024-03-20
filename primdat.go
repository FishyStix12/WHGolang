///////////////////////////////////////////////////////////////////////////////////////////////////////////
// Author: Nicholas Fisher
// Date: March 18th 2024
// Description of the Script
// The provided Go script initializes two variables: `x` with the string "Hello New World" and `z` with 
// the integer 42. It then prints the values of these variables to the console.
// Primative Data Type:
// A primitive data type is a basic building block of a programming language that represents simple
// values, such as numbers (like integers and floating-point numbers) or characters (like letters and 
// symbols). These data types are predefined by the programming language and are used to store and manipulate
// data in programs. Primitive data types typically have a fixed size and behavior defined by the programming
// language. In simple terms, you can think of them as the elemental pieces of data that a program can work 
// with, like individual bricks in a construction project.
///////////////////////////////////////////////////////////////////////////////////////////////////////////
package main // Declares the package name. A Go program starts running in package main.

import "fmt" // Imports the fmt package, which implements formatted I/O.

func main() { // Declares the main function, which serves as the entry point of the program.
	var x = "Hello New World" // Declares a variable named x of type string and assigns it the value "Hello New World".

	z := int(42) // Declares a variable named z and initializes it with the integer value 42.
	// The := operator is a shorthand for variable declaration and initialization.

	fmt.Println(x) // Prints the value of x to the standard output using the fmt.Println function.
	fmt.Println(z) // Prints the value of z to the standard output using the fmt.Println function.
}
