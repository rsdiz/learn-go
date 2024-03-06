package main

import (
	"fmt"
	"reflect"
)

func main() {
	/**
	in go-lang, we can convert a data type to another data type
	but please note,to perform data type conversions carefully
	because when value does not fit within its allocated memory space,
	an overflow will occur
	Example: Integer Overflow will occur when 32768 convert to int8
	*/

	var data1 int32 = 129
	var data2 = int16(data1)
	var data3 = int8(data2) // occur integer overflow

	fmt.Println(data1)
	fmt.Println(data2)
	fmt.Println(data3)

	var name = "Rosyid"
	var r = name[0] // by default, this will return uint8 data type
	var rString = string(r)

	fmt.Println(name)
	fmt.Println(r)
	fmt.Println(rString)

	/**
	to check the data type from variable, we can use library from reflect.
	reflect.TypeOf(variableName)
	*/

	fmt.Println("data1 is", reflect.TypeOf(data1))
	fmt.Println("data2 is", reflect.TypeOf(data2))
	fmt.Println("data3 is", reflect.TypeOf(data3))
	fmt.Println("name is", reflect.TypeOf(name))
	fmt.Println("r is", reflect.TypeOf(r))
	fmt.Println("rString is", reflect.TypeOf(rString))
}
