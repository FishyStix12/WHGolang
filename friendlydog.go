///////////////////////////////////////////////////////////////////////////////////////////////////////////
// Author: Nicholas Fisher
// Date: March 18th 2024
// Description of the Script
//  The provided script is a concise Go program that demonstrates object-oriented principles such as structs,
// methods, and interfaces. It defines a `Person` struct with a `SayHello()` method, as well as an `interface`
// named `Friend` with the same method signature. Additionally, it defines a `Dog` struct with its own 
// `SayHello()` method. The `main()` function instantiates a `Person` and a `Dog`, calls their respective 
// `SayHello()` methods, and utilizes the `Greet()` function to greet both a person and a dog interchangeably 
// through the `Friend` interface. This showcases polymorphism in Go, where different types can be treated 
// uniformly through interfaces.
///////////////////////////////////////////////////////////////////////////////////////////////////////////
// Declares the package name. A Go program starts running in package main.
package main

// Imports the fmt package, which implements formatted I/O.
import "fmt"

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

// Defines an interface named Friend with a single method SayHello.
type Friend interface {
	SayHello()
}

// Function to greet a Friend interface.
func Greet(f Friend) {
	f.SayHello()
}

// Defines a struct named Dog with no fields.
type Dog struct{}

// Defines a method named SayHello associated with the Dog struct.
// This method prints "Bark Bark".
func (d *Dog) SayHello() {
	fmt.Println("Bark Bark")
}

// Main function, the entry point of the program.
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

	// Calls the Greet() function with guy as an argument.
	// Since guy satisfies the Friend interface, it calls the SayHello() method of Person.
	Greet(guy)

	// Creates a new instance of the Dog struct using the new() function.
	// Assigns it to the variable dog.
	var dog = new(Dog)

	// Calls the Greet() function with dog as an argument.
	// Since dog satisfies the Friend interface, it calls the SayHello() method of Dog.
	Greet(dog)
}
