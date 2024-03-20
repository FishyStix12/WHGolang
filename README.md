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

**3. Pointers** <br />
     A pointer is a variable that stores the memory address of another variable. Pointers allow indirect access to the value of the variable they point to. They are used to share data between different parts of a program efficiently and to enable the manipulation of data indirectly. In Go, pointers are represented using the * symbol followed by the type of the variable they point to. <br />

**4. Struct** <br />
     A struct is a composite data type that groups together zero or more fields of possibly different types under a single name. It allows you to create custom data types by defining the structure of the data, similar to classes in object-oriented programming languages. Each field within a struct can have its own name and type. Structs are used to represent more complex data structures and enable better organization and management of related data in your programs. <br />
     
**5. Interface** <br />
     An interface is a type that specifies a set of method signatures. An interface defines behavior, but it does not provide implementations for the methods it declares. Instead, any type that implements all the methods declared in the interface is said to satisfy or implement the interface implicitly. This allows for polymorphic behavior, where different types can be treated interchangeably if they implement the same interface. Interfaces promote loose coupling between components and enable abstraction, making it easier to write modular and extensible code in Go. <br />

**The Following List gives a short description of all the scripts in this group:** <br />
1. hello_world.go - The script is a simple Go program that demonstrates the basic structure of a Go application. It starts by declaring that it belongs to the `main` package, imports the "fmt" package for formatting input and output, defines the `main` function as the entry point of the program, and then prints "Hello, White Hat Gophers!" to the console using the `fmt.Println` function. This script serves as a starting point for learning how to write and execute Go programs. <br />
2. primdat.go - The provided Go script initializes two variables: `x` with the string "Hello New World" and `z` with the integer 42. It then prints the values of these variables to the console. A primitive data type is a basic building block of a programming language that represents simple values, such as numbers (like integers and floating-point numbers) or characters (like letters and symbols). These data types are predefined by the programming language and are used to store and manipulate data in programs. Primitive data types typically have a fixed size and behavior defined by the programming language. In simple terms, you can think of them as the elemental pieces of data that a program can work with, like individual bricks in a construction project. <br />
3. slicemaps.go - The code provided is a simple Go program that demonstrates the use of slices and maps, which are two common complex data types in Go. <br />
pointer.go - The provided Go script demonstrates the usage of pointers in Go programming language. It initializes an integer variable `count` with the value 42, creates a pointer `ptr` to this variable, and then prints the value of `count` via the pointer. Afterward, it modifies the value of `count` through the pointer, assigning it the value 100, and prints the updated value of `count`. This script illustrates how pointers can be used to indirectly access and modify variables in Go. <br />
4. struct.go - This Go code defines a simple program that demonstrates the use of structs, methods, and basic I/O in Go. It declares a struct type Person with fields for name and age, and a method SayHello() to greet a person by name. In the main() function, it creates an instance of Person, sets the name, and calls the SayHello() method to print a greeting. <br />
5. interface.go - This Go script defines a Person struct with fields for name and age, along with a method SayHello() to greet a person by printing their name. It also defines an interface Friend with a method SayHello(). The Person struct implicitly implements this interface. In the main() function, it creates a person named "Nicholas Fisher" and greets them using the SayHello() method. <br />
6. function.go - This Go script defines a Person struct with a SayHello() method and an interface Friend with a SayHello() method. It demonstrates creating an instance of the Person struct, setting its name, and calling the SayHello() method. Additionally, it showcases the use of interfaces in Go. <br />

**Example output of some of the go commands, and scripts:** <br />
1. "go doc fmt.Println" command output: <br />
    func Println(a ...any) (n int, err error) <br />
         Println formats using the default formats for its operands and writes to <br />
         standard output. Spaces are always added between operands and a newline <br />
         is appended. It returns the number of bytes written and any write error <br />
         encountered. <br />
