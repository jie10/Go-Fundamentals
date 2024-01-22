// Package variables Declare the package `variables` is in. This is the entry point for the program execution
package main

// Import required package fmt, which includes functions for printing output to console among others
import "fmt"

// Declare a global variable `globalVar` of type string, accessible throughout this package
var globalVar = "This is a global variable"

// Declare a constant `name` of type string, which cannot be changed after declaration
const name = "Hello world"

// Define the main function. This function gets executed when the program runs.
func main() {
	// Declare a local variable `version` of type integer. Its scope is limited to the `main` function
	version := 1 // type is inferred as int

	// Print the local variable `version` on the console
	fmt.Println(version)

	// Print the constant variable `name` on the console
	fmt.Println(name)

	// Print the global variable `globalVar` on the console
	fmt.Println(globalVar)
}
