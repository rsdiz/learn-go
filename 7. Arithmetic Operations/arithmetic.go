package main

import "fmt"

func main() {
	/**
	Arithmetic operation in go-lang is same as other programming language
	+ (Add)
	- (Reduction)
	* (Multiply)
	/ (Divider)
	% (Modulo)
	*/

	sum := 10 + 10
	fmt.Println("Penjumlahan:", sum)
	sum = 10 - 10
	fmt.Println("Pengurangan:", sum)
	sum = 10 * 10
	fmt.Println("Perkalian:", sum)
	sum = 10 / 10
	fmt.Println("Pembagian:", sum)
	sum = 10 % 10
	fmt.Println("Modulus:", sum)

	/**
	Like in other programming language, go-lang has augmented assignments too
	+=
	-=
	*=
	/=
	%=
	*/

	angka := float32(1)
	fmt.Println("Angka awal:", angka)
	angka += 9
	fmt.Println("Angka + 9:", angka)
	angka -= 4
	fmt.Println("Angka - 4:", angka)
	angka *= 3
	fmt.Println("Angka * 3:", angka)
	angka /= 5.0
	fmt.Println("Angka / 5:", angka)
	intAngka := int(angka)
	intAngka %= 2
	fmt.Println("Angka % 2:", intAngka)

	/**
	Then, another operation is unary operator
	++ (increment)
	-- (decrement)
	- (negative)
	+ (positive)
	! (Reverse Boolean)
	*/

	count := 1
	fmt.Println("Initialize:", count)
	count++
	fmt.Println("Increment:", count)
	count--
	fmt.Println("Decrement:", count)
	count = -count
	fmt.Println("Negative:", count)
	state := true
	fmt.Println("Current State:", state)
	fmt.Println("Change State:", !state)
}
