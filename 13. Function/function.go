package main

import (
	"fmt"
	"strings"
)

type (
	Filter    func(string) string
	Blacklist func(string) bool
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

func getHelloWithFilter(word string, filter Filter) string {
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

func login(username string, blacklist Blacklist) {
	if blacklist(username) {
		fmt.Println("Cannot Logged In! You're blocked by System!")
	} else {
		fmt.Println("Welcome Back", username)
	}
}

func factorialWithLoop(number int) (result int) {
	result = 1
	for i := number; i > 0; i-- {
		result *= i
	}
	return
}

func factorialWithRecursive(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * factorialWithRecursive(number-1)
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

	// Function Type Declarations
	// Sometimes function is too long, too hard to write function as parameter,
	// We can create alias function with Type Declarations, so it will make it easy for us
	// when We use function as parameter

	// see parameter in getHelloWithFilter(), now it using type declaration
	filteredHello = getHelloWithFilter("Celeng", filterBadWord)
	fmt.Println(filteredHello)

	// Anonymous Function
	// Previous, everytime We create function, We're gives name to it
	// But sometimes, for make it easier we store function in variable or as parameter without create the function first.
	// That is named Anonymous Function, or function without name

	// First Method, store as variable
	blacklist := func(username string) bool {
		switch strings.ToLower(username) {
		case "root":
			return true
		case "admin":
			return true
		default:
			return false
		}
	}
	login("rosyid", blacklist)

	// Second Method, write directly as parameter when calling function
	login("Hacker", func(username string) bool {
		return strings.ToLower(username) == "hacker"
	})

	// Recursive Function
	// is function that calls the function itself
	// Sometimes we need to write recursive function, because it easier than when We write without it.
	// Example case is write function to calculate factorial number

	fmt.Println(4 * 3 * 2 * 1)             // This is how factorial works
	fmt.Println(factorialWithLoop(4))      // This is example without recursive
	fmt.Println(factorialWithRecursive(4)) // This is example with recursive

	// Closures
	// is an ability of function to interact with data's around it in the same scope

	counter := 0
	plusWith := func(value int) {
		// variable counter in above function can be accessed in this function
		// but variable in this scope can't be accessible outside this function
		counter += value
	}
	plusWith(2)
	plusWith(4)
	fmt.Println("Current Counter:", counter)
}
