package main

import "fmt"

func main() {
	// Boolean operation:
	// && (and)
	// || (or)
	// ! (not)

	finalScore := 87
	presence := 90

	passedFinalScore := finalScore > 80
	passedPresence := presence > 80

	passed := passedFinalScore && passedPresence

	fmt.Println("Passed:", passed)
}
