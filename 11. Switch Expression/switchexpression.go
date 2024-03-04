package main

import "fmt"

func main() {
	// Normal Switch

	age := 20
	job := "Programmer"

	switch age {
	case 20:
		fmt.Println("Age is 20")
		fmt.Println("Your job is", job)
	case 10:
		fmt.Println("Age not met to have a job!")
	default:
		fmt.Println("Age is not met requirement!")
	}

	// Switch with Short Statement

	switch length := len(job); length < 10 && age >= 20 {
	case true:
		fmt.Println("Job length is short")
	case false:
		fmt.Println("Job length is normal")
	}

	// Switch without Condition

	switch {
	case age > 10 && age <= 17:
		fmt.Println("Please go to school!")
	case age >= 20:
		fmt.Println("Please go to work!")
	default:
		fmt.Println("You are to young!")
	}
}
