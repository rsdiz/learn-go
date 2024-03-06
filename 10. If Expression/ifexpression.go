package main

import "fmt"

func main() {
	/**
	IF,
	If is one of the keyword in conditional statement
	if block will be executed when the condition is true

	ELSE,
	Else block will be executed when the condition are not met/false

	ELSE-IF,
	Else-if block will be executed when previous condition are false, and current if is true
	*/

	currentDay := "Senin"

	if currentDay == "Senin" {
		fmt.Println("Semangat Pejuang Senin!")
	} else if currentDay == "Selasa" {
		fmt.Println("Tetap Semangat walaupun Sekarang Hari Selasa!")
	} else {
		fmt.Println("Sekarang bukan Hari Senin maupun hari Selasa!")
	}

	/**
	If with Short Statement
	in Go-Lang we can create if statement with short way
	*/

	if count := len(currentDay); count > 5 {
		fmt.Println("Nama hari lebih dari 5 huruf!")
	} else {
		fmt.Println("Nama harinya singkat!")
	}
}
