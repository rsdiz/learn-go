package main

import "fmt"

func PrintHello() {
	fmt.Println("Hello")
}

func PrintHelloTo(name string) {
	fmt.Println("Hello", name)
}

func getHello(name string) string {
	return "Hello again " + name
}

func main() {
	// Function
	// is a block of code intentionally created in a program so that can be used repeatedly
	// to create function, we need keyword "func" following by the function name, and the code block

	PrintHello()

	// Parameter Function
	// When we create function, sometimes we need data from outside, or what we call parameters
	// We can input parameter more than one
	// Parameter is not required, so we can create function without parameter

	name := "Rosyid"
	PrintHelloTo(name)

	// Function Return Value
	// Function can return a value
	// To create function with return value, we need to write down the return data type of the function

	fmt.Println(getHello(name))
}
