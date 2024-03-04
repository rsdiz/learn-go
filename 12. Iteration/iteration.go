package main

import "fmt"

func Repeat(char string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += char
	}
	return repeated
}

func main() {
	// For loops is one of the iteration
	// Simple For
	counter := 1

	for counter <= 10 {
		fmt.Println("Loop!", counter)
		counter++
	}

	// For with Statement
	for index := 0; index < 5; index++ { // for init statement; statement; post statement {...}
		fmt.Println("Loop!", index)
	}

	// For Range
	// can be used for array, slice, and map
	data := []string{
		"Muhammad",
		"Rosyid",
		"Izzulkhaq",
	}

	for index, value := range data {
		fmt.Printf("Index: %d = %q\n", index, value)
	}

	fmt.Println(Repeat("F", 5))
}
