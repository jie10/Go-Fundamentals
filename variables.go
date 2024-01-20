package main

import "fmt"

// Global Variables
var globalVar = "This is global var"

// Constant
const name = "Hello world"

func main() {
	// Local Variables
	version := 1 // infer int

	// print local var
	fmt.Println(version)

	// print global var
	fmt.Println(globalVar)
}
