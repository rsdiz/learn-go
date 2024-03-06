package main

import "fmt"

func main() {
	/**
	Type declarations is ability to recreate new data types from existing data types
	Type declarations usually used to create alias
	To create type declaration, use keyword:
	type typeName dataType
	*/

	type Cuaca string

	var overcast Cuaca = "Mendung"
	rain := Cuaca("Hujan")

	fmt.Println("Suasana saat ini", overcast)
	fmt.Println("Kemungkinan nanti akan", rain)
}
