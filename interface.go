///////////////////////////////////////////////////////////////////////////////////////////////////////////
// Author: Nicholas Fisher
// Date: March 18th 2024
// Description of the Script
//  This Go script defines a Person struct with fields for name and age, along with a method SayHello()
// to greet a person by printing their name. It also defines an interface Friend with a method SayHello().
// The Person struct implicitly implements this interface. In the main() function, it creates a person named
// "Nicholas Fisher" and greets them using the SayHello() method.
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

// Defines an interface named Friend with a single method SayHello.
type Friend interface {
    SayHello()
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
