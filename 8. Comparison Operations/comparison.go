package main

import "fmt"

func main() {
	/**
	Comparison operation is operation for compare two data
	the operation will produce a boolean value
	operator:
	> (more than)
	< (less than)
	>= (more than equal)
	<= (less than equal)
	== (equal)
	!= (not equal to)
	*/

	value1 := 100
	value2 := 150

	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 >= value2)
	fmt.Println(value1 <= value2)
}
