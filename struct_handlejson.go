///////////////////////////////////////////////////////////////////////////////////////////////////////////
// Author: Nicholas Fisher
// Date: March 18th 2024
// Description of the Script
// This Go script demonstrates basic JSON encoding and decoding operations. It defines a struct `Foo` with 
// two string fields, `Bar` and `Baz`. In the `main()` function, it initializes an instance of `Foo` with 
// specific values. The script then encodes this instance into a JSON string using `json.Marshal()` and 
// prints it out. Afterward, it decodes the JSON string back into a `Foo` instance using `json.Unmarshal()`,
// demonstrating the reverse process. However, the script lacks error handling for the decoding operation. 
// Overall, it serves as a simple illustration of how to use Go's standard library for JSON serialization
// and deserialization.
///////////////////////////////////////////////////////////////////////////////////////////////////////////
package main // Declares the package name. A Go program starts running in package main.

import (
	"encoding/json" // Importing the encoding/json package for JSON encoding and decoding.
	"fmt"           // Importing the fmt package for formatted I/O.
)

// Defining a struct named Foo with two string fields: Bar and Baz.
type Foo struct {
	Bar string
	Baz string
}

func main() {
	// Initializing an instance of Foo with specific values.
	f := Foo{"Nicholas Fisher", "Hello FishyStix12"}

	// Encoding the Foo instance f into JSON format and storing it in byte slice b.
	b, _ := json.Marshal(f)

	// Printing the JSON string representation of the Foo instance.
	fmt.Println(string(b))

	// Decoding the JSON string stored in byte slice b back into the Foo instance f.
	json.Unmarshal(b, &f)
}
