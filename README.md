# WHGolang <br />
# By: Nicholas Fisher <br />
**Disclaimer: These scripts should only be used in cases where the user has permission to use these scripts on the subject systems!** <br />

**Import Note: Please read this statement carefully: By downloading any of the scripts in this repository, you, as the user, take full responsibility for storing, and using these scripts. You also take full responsibility for any misuse of these malicious codes.** <br />

# Golang Linux Setup <br />
![image](https://github.com/FishyStix12/White-Hat-Go-/assets/102126354/05e0dca4-a04f-4ed1-842e-5ab143ace30a) <br />
**Please run these scripts in your home directory!** <br />

**The Following List gives a short description of all the scripts in this group:** <br />
1. goconfig.sh - This Bash script installs the Go programming language on a Linux system using the `sudo apt install golang-go` command. It then sets up environment variables (`GOROOT`, `GOPATH`, and updates `PATH`) in the `~/.bashrc` file to configure the Go development environment. Finally, it reloads the `~/.bashrc` file using the `source ~/.bashrc` command to apply the changes immediately. <br />
2. gomods.sh - This Bash script installs the latest versions of several Go programs from their respective GitHub repositories. The programs being installed are: <br />
               1. dalfox by hahwul for web application vulnerability scanning. <br />
               2. qsreplace by tomnomnom for URL query string manipulation. <br />
               3. hakrawler by hakluke for website crawling and analysis. <br />
               4. httprobe by tomnomnom for HTTP/HTTPS probe. <br />
               5. wildcheck by theblackturtle for wildcard domain detection. <br />
               6. gau by lc for Google Analytics data collection. <br />
               7. assetfinder by tomnomnom for finding domains and subdomains. <br />

# Fundamentals of Golang <br />
![image](https://github.com/FishyStix12/White-Hat-Go-/assets/102126354/4650c928-f66a-494b-9c05-e02af87698e9) <br />
**The "go doc \<command\>" pulls documentation about a function, package, variable, or method, and is embedded as comments in the code!** <br />

**The "go fmt \/path\/to\/your\/script" command automatically formats your source code!** <br />

**The "golint \/path\/to\/your\/script" command reports style mistakes such as missing comments, variables names that don't follow conventions, etc.** <br />

**Install goline by entering "go get -u golang.org\/x\/lint\/golint"** <br />

**The "go vet \/path\/to\/your\/script" command attempts to identify issues, some of which might be bugs, that the compiler might miss!** <br />

**Explanation of the Structure of Golang:** <br />
**1. Slices:** <br />
     A slice in Go is a data structure that provides a flexible way to work with sequences of elements. It's like a dynamically-sized array. Slices are created using the make function, as shown in the code. They are defined by specifying the type of elements they contain (in this case, strings). Slices can grow and shrink dynamically, making them useful for managing collections of data. <br />

**2. Maps:** <br />
     A map in Go is a data structure that represents a collection of key-value pairs, where each key is unique. It's similar to a dictionary in Python or an associative array in other languages. Maps are created using the make function, as demonstrated in the code. In the code, a map with string keys and string values is declared. Values can be accessed and modified using their corresponding keys. <br />

**3. Pointers:** <br />
     A pointer is a variable that stores the memory address of another variable. Pointers allow indirect access to the value of the variable they point to. They are used to share data between different parts of a program efficiently and to enable the manipulation of data indirectly. In Go, pointers are represented using the * symbol followed by the type of the variable they point to. <br />

**4. Struct:** <br />
     A struct is a composite data type that groups together zero or more fields of possibly different types under a single name. It allows you to create custom data types by defining the structure of the data, similar to classes in object-oriented programming languages. Each field within a struct can have its own name and type. Structs are used to represent more complex data structures and enable better organization and management of related data in your programs. <br />
     
**5. Interface:** <br />
     An interface is a type that specifies a set of method signatures. An interface defines behavior, but it does not provide implementations for the methods it declares. Instead, any type that implements all the methods declared in the interface is said to satisfy or implement the interface implicitly. This allows for polymorphic behavior, where different types can be treated interchangeably if they implement the same interface. Interfaces promote loose coupling between components and enable abstraction, making it easier to write modular and extensible code in Go. <br />

**6. Function:** <br />
     A function is a reusable block of code that performs a specific task. It consists of a function signature and a function body. The signature includes the function's name, parameters (if any), and return type (if any). The function body contains the code that defines the functionality of the function. <br />

**7. Switch:** <br />
     A switch statement is used to evaluate an expression and match it against multiple possible cases. It provides a concise way to express a multi-way branch. <br />

**8. For Loop:** <br />
     A for loop is used to repeatedly execute a block of code until a specified condition is met. <br />

**9. GoRoutines:** <br />
     Goroutines are lightweight threads of execution managed by the Go runtime. Goroutines allow functions to be executed concurrently with other functions. They are multiplexed onto multiple OS threads so that multiple goroutines can run concurrently within a single OS thread. Goroutines are created using the go keyword followed by a function call. They provide a simple and efficient mechanism for concurrent programming in Go. <br />

**10. Channels:** <br />
      Channels are a way for goroutines to communicate with each other and synchronize their execution. A channel in Go is a typed conduit through which you can send and receive values with the channel operator, <-. Channels provide a way for two goroutines to synchronize execution and coordinate their work. They can be thought of as pipes through which data can flow. Channels can be either buffered or unbuffered, and they enforce communication patterns between goroutines. <br />

**The Following List gives a short description of all the scripts in this group:** <br />
1. hello_world.go - The script is a simple Go program that demonstrates the basic structure of a Go application. It starts by declaring that it belongs to the `main` package, imports the "fmt" package for formatting input and output, defines the `main` function as the entry point of the program, and then prints "Hello, White Hat Gophers!" to the console using the `fmt.Println` function. This script serves as a starting point for learning how to write and execute Go programs. <br />
2. primdat.go - The provided Go script initializes two variables: `x` with the string "Hello New World" and `z` with the integer 42. It then prints the values of these variables to the console. A primitive data type is a basic building block of a programming language that represents simple values, such as numbers (like integers and floating-point numbers) or characters (like letters and symbols). These data types are predefined by the programming language and are used to store and manipulate data in programs. Primitive data types typically have a fixed size and behavior defined by the programming language. In simple terms, you can think of them as the elemental pieces of data that a program can work with, like individual bricks in a construction project. <br />
3. slicemaps.go - The code provided is a simple Go program that demonstrates the use of slices and maps, which are two common complex data types in Go. <br />
pointer.go - The provided Go script demonstrates the usage of pointers in Go programming language. It initializes an integer variable `count` with the value 42, creates a pointer `ptr` to this variable, and then prints the value of `count` via the pointer. Afterward, it modifies the value of `count` through the pointer, assigning it the value 100, and prints the updated value of `count`. This script illustrates how pointers can be used to indirectly access and modify variables in Go. <br />
4. struct.go - This Go code defines a simple program that demonstrates the use of structs, methods, and basic I/O in Go. It declares a struct type Person with fields for name and age, and a method SayHello() to greet a person by name. In the main() function, it creates an instance of Person, sets the name, and calls the SayHello() method to print a greeting. <br />
5. interface.go - This Go script defines a Person struct with fields for name and age, along with a method SayHello() to greet a person by printing their name. It also defines an interface Friend with a method SayHello(). The Person struct implicitly implements this interface. In the main() function, it creates a person named "Nicholas Fisher" and greets them using the SayHello() method. <br />
6. function.go - This Go script defines a Person struct with a SayHello() method and an interface Friend with a SayHello() method. It demonstrates creating an instance of the Person struct, setting its name, and calling the SayHello() method. Additionally, it showcases the use of interfaces in Go. <br />
7. friendlydog.go - The provided script is a concise Go program that demonstrates object-oriented principles such as structs, methods, and interfaces. It defines a `Person` struct with a `SayHello()` method, as well as an `interface` named `Friend` with the same method signature. Additionally, it defines a `Dog` struct with its own `SayHello()` method. The `main()` function instantiates a `Person` and a `Dog`, calls their respective `SayHello()` methods, and utilizes the `Greet()` function to greet both a person and a dog interchangeably through the `Friend` interface. This showcases polymorphism in Go, where different types can be treated uniformly through interfaces. <br />
8. contstruct.go - The script is a simple Go program that demonstrates the usage of conditionals (if statement) and switch statements. It defines a variable x, assigns it the value "foo", and then checks whether x is equal to "1" using an if statement. Depending on the value of x, it prints out a corresponding message. Additionally, it employs a switch statement to determine the action based on the value of x, printing different messages for different cases. <br />
9. type_printer.go - The code in `type_printer.go` is a simple Go script that defines a function called `foo` which takes an interface argument and prints a message indicating the type of the argument. It utilizes a type switch to handle different types of input: integers, strings, and booleans. The `main` function demonstrates the usage of the `foo` function with various types of inputs and prints out the corresponding type messages. This script serves as a basic illustration of type assertion and type switching in Go. <br />
10. forloop.go - The provided Go code demonstrates the usage of a for loop and range iteration in Go. It first prints numbers from 0 to 10 using a simple for loop, then it iterates over the elements of a slice (nums) using the range keyword, printing the index and value of each element. <br />
11. goroutines.go - The script demonstrates the use of goroutines and channels in Go. It defines a function f() that prints a message, and a function strlen() that calculates the length of a string and sends it to a channel. In the main() function, f() is executed concurrently in a goroutine, followed by a 1-second sleep. Then, the length of the string "Salutations" is calculated concurrently using strlen() and sent to a channel. The main function receives two values from the channel (the length of "Salutations" and 0), prints them, and prints their sum. <br />
12. channels.go - This script demonstrates a simple concurrency example in Go. It defines a function strlen which calculates the length of a given string and sends it to a channel. Then, in the main function, it creates a channel, launches strlen in a separate goroutine, receives the length of the string from the channel, and prints the sum of the lengths of the string. <br />
13. error_handle.go - The code defines a custom error type MyError, implements the error interface for it, and creates a function foo() that returns an error of this custom type. In the main() function, it calls foo() and checks if an error is returned. If an error occurs, it prints out the error message. <br />
14. struct_handlejson.go - This Go script demonstrates basic JSON encoding and decoding operations. It defines a struct `Foo` with two string fields, `Bar` and `Baz`. In the `main()` function, it initializes an instance of `Foo` with specific values. The script then encodes this instance into a JSON string using `json.Marshal()` and prints it out. Afterward, it decodes the JSON string back into a `Foo` instance using `json.Unmarshal()`, demonstrating the reverse process. However, the script lacks error handling for the decoding operation. Overall, it serves as a simple illustration of how to use Go's standard library for JSON serialization and deserialization. <br />

**Example output of some of the go commands, and scripts:** <br />
1. "go doc fmt.Println" command output: <br />
    func Println(a ...any) (n int, err error) <br />
         Println formats using the default formats for its operands and writes to <br />
         standard output. Spaces are always added between operands and a newline <br />
         is appended. It returns the number of bytes written and any write error <br />
         encountered. <br />

# Network Tools <br />
![image](https://github.com/FishyStix12/WHGolang/assets/102126354/8b870a1e-d488-4851-8f94-74511c35c748) <br />
**The Following List gives a short description of all the scripts in this group:** <br />
1. SimpleScan.go - This Go script attempts to establish a TCP connection to the server "scanme.nmap.org" on port 80. If the connection is successful, it prints "Connection successful" to the console. <br />
2. Concurrentscan.go - This Go script conducts a TCP port scan on the server "scanme.nmap.org" across ports 1 through 1024. It utilizes concurrency to efficiently scan multiple ports simultaneously. For each port, it attempts to establish a TCP connection. If the connection is successful, it indicates that the port is open; otherwise, it assumes the port is either closed or filtered. The script prints the status of each port, either indicating it as open or specifying that it is closed or filtered. <br />
3. toofast.go - This Go script is designed to quickly scan for open TCP ports on the host `scanme.nmap.org`. It targets ports from 1 to 1024, using goroutines to perform the scan concurrently, which significantly speeds up the process. For each port, the script tries to establish a TCP connection. If the connection succeeds, the port is marked as open, and its number is printed to the console. Each connection attempt runs in its own goroutine, allowing the script to efficiently complete the port scan. To test the speed of the script, you can run it and measure the time it takes with the command `time ./toofast.go`. <br />
4. worksfast.go - The provided Go script is a concurrent port scanner that checks the availability of ports from 1 to 1024 on the host `scanme.nmap.org`. It utilizes goroutines to handle each port scan concurrently, significantly speeding up the process. A `sync.WaitGroup` is used to ensure the main function waits for all goroutines to complete before exiting. For each port, the script attempts to establish a TCP connection; if successful, it prints the port number and closes the connection. Errors, such as connection refusals, are handled gracefully by simply returning from the goroutine. This approach demonstrates efficient use of Go's concurrency features to perform network operations. <br />

**Example output of some of the go commands, and scripts:** <br />
1. Concurrentscan.go: <br />
   Port 1 is closed or filtered <br />
   2 is open! <br />
   Port 3 is closed or filtered <br />
   Port 4 is closed or filtered <br />
   ... <br />
   80 is open! <br />
   Port 81 is closed or filtered <br />
   ... <br />
   443 is open! <br />
   Port 444 is closed or filtered <br />
   ... <br />
   1023 is open! <br />
   Port 1024 is closed or filtered <br />
2. time ./toofast.go: <br />
   real    0m1.123s <br />
   user    0m0.567s <br />
   sys     0m0.234s <br />
