///////////////////////////////////////////////////////////////////////////////////////////////////////////
// Author: Nicholas Fisher
// Date: March 18th 2024
// Description of the Script
//  This Go code defines a simple program that demonstrates the use of structs, methods, and basic 
// I/O in Go. It declares a struct type Person with fields for name and age, and a method SayHello() to 
// greet a person by name. In the main() function, it creates an instance of Person, sets the name, 
// and calls the SayHello() method to print a greeting.
///////////////////////////////////////////////////////////////////////////////////////////////////////////
package main // Declares the package name. A Go program starts running in package main.

import "fmt" // Imports the fmt package, which implements formatted I/O.

// Defines a struct named Person with two fields: Name (string) and Age (int).
type Person struct {
	Name string
	Age  int
}

// Defines a method named SayHello associated with the Person struct.
// This method takes a pointer to a Person (*Person) as its receiver.
// It prints "Hello, " followed by the name of the person.
func (p *Person) SayHello() {
	fmt.Println("Hello, ", p.Name)
}

func main() {
	// Declares a variable named guy of type *Person.
	// Initializes it with a new instance of the Person struct using the new() function.
	// This allocates memory for a Person struct and returns a pointer to it.
	var guy = new(Person)

	// Sets the Name field of the guy struct to "Nicholas Fisher".
	guy.Name = "Nicholas Fisher"

	// Calls the SayHello() method on the guy struct pointer.
	// This prints "Hello, Nicholas Fisher" to the console.
	guy.SayHello()
}
