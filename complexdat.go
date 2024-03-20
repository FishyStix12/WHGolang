///////////////////////////////////////////////////////////////////////////////////////////////////////////
// Author: Nicholas Fisher
// Date: March 18th 2024
// Description of the Script
//  The code provided is a simple Go program that demonstrates the use of slices and maps, which are 
// two common complex data types in Go.
///////////////////////////////////////////////////////////////////////////////////////////////////////////
package main // Declares the package name. A Go program starts running in package main.

import "fmt" // Imports the fmt package, which implements formatted I/O.

func main() { // Declares the main function, which serves as the entry point of the program.
    // Declare an empty slice of strings
    s := make([]string, 0)

    // Declare an empty map with keys and values of type string
    m := make(map[string]string)

    // Append a string "Example of" to the slice
    s = append(s, "Example of")

    // Assign the value "Data Types" to the key "Complex" in the map
    m["Complex"] = "Data Types"

    // Print the contents of the slice
    fmt.Println("Slice:")
    for _, val := range s { // Iterate over the slice 's', '_' is used to ignore the index value
        fmt.Println(val) // Print each element of the slice
    }

    // Print the contents of the map
    fmt.Println("\nMap:")
    for key, val := range m { // Iterate over the map 'm', 'key' and 'val' represent the key-value pair
        fmt.Printf("%s: %s\n", key, val) // Print each key-value pair of the map
    }
}
