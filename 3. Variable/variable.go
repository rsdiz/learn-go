package main

import "fmt"

func main() {
	// in golang, variable only store the same data type
	// if we want to store data of different types, we have to create several variables
	// to create variable, use keyword "var" follow with variable name and data type

	var name string
	name = "Rosyid Izzulkhaq"
	fmt.Println("Nama:", name)

	name = "Muhammad Rosyid Izzulkhaq"
	fmt.Println("Update Nama:", name)

	// Data type is important went we create the variable,
	// but if we create variable, and we initialize it with some value,
	// we don't need to put the data type

	var long = 3
	fmt.Println("How long:", long)

	long = 7
	fmt.Println("Updated long:", long)

	// instead using "var" to create variable, we can use ":=" when we initialize variable

	age := 24
	fmt.Println("Umur Saya:", age)

	age = 25
	fmt.Println("Umur Saya Tahun Depan:", age)

	// in golang we can declare multiple variable at the same time

	var (
		firstName = "Muhammad"
		lastName  = "Izzulkhaq"
	)
	fmt.Println("Nama Depan:", firstName, "\nNama Belakang:", lastName)
}
