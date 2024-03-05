package main

import (
	"fmt"
	"strings"
)

func PrintHello() {
	fmt.Println("Hello")
}

func PrintHelloTo(name string) {
	fmt.Println("Hello", name)
}

func getHello(name string) string {
	return "Hello again " + name
}

func getFullName() (string, string) {
	return "Rosyid", "Izzulkhaq"
}

func getAccessoriesStats() (bracelet, necklace, earring int) {
	bracelet = 25
	necklace = 30
	earring = 32

	return
}

func calcAverage(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total / len(numbers)
}

func getHelloWithFilter(word string, filter func(string) string) string {
	return getHello(filter(word))
}

func filterBadWord(word string) string {
	switch strings.ToLower(word) {
	case "asu":
		return "..."
	case "celeng":
		return "..."
	default:
		return word
	}
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

	// Returning Multiple Values
	// in Go-Lang, function can return multiple values
	// to create function with return multiple values, we need to write down all the return data type of the function

	firstName, lastName := getFullName()
	fmt.Println("Firstname:", firstName, "\nLastname:", lastName)

	// Ignore Return Value
	// We can ignore return value by change the variable to underscore (_)

	firstName, _ = getFullName()

	// Named Return Value
	// in Go-Lang we can give name to return value

	bracelet, necklace, earring := getAccessoriesStats()
	fmt.Printf("Bracelet give extra %d ATK, Necklace give extra %d DEF, and Earring give extra %d SPD\n", bracelet, necklace, earring)

	// Variadic Function
	// The last parameter in function has an ability to become a varargs.
	// Varargs means the data can accept more than one input, or it as some kind of array

	average := calcAverage(100, 96, 80, 85, 92, 83)
	fmt.Println("Average grade from Class A is", average)

	// Variadic Function but Slice as Parameter
	// Yes, We can make vararg with slice as parameter

	grades := []int{100, 96, 80, 85, 92, 83}
	average = calcAverage(grades...)
	fmt.Println("Average grade from Class A is", average)

	// Function Value
	// Function is first class citizen, function can create as datatype, and can be stored in variable
	// It means, we can create function as independent without write class, like in OOP

	hello := getHello // store function in variable
	fmt.Println(hello("Rosyid"))

	// Function as Parameter
	// Besides being able to store in variable as value, we can also use function as parameters for other functions

	filteredHello := getHelloWithFilter("ASU", filterBadWord)
	fmt.Println(filteredHello)
}
