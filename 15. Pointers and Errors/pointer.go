package main

import (
	"errors"
	"fmt"
)

/**
Pointer.
By default, all variable in go-lang is passing by value, not by reference,
it's mean, if we send variable in the function, method, or other variable,
what is actually sent is a duplicate value.

Pointer is ability to make references to data locations in the same memory, without duplicating existing data.
It's mean, by using pointer, we can create pass by reference.

Operator &
To create a variable with pointer value to other variable, we can use "&" operator, following by variable name.

Operator *
When we update pointer variable, then the only things that change is that variable,
all variables that refer to the same data will not change.
If we want to change all reference variable, we can use "*" operator.

Function new
Previous, to create pointer we need "&" operator.
Go-lang has a function to create new pointer, it's called new(),
but this function just return empty data, without initial data
*/

type (
	Rupiah int

	Person struct {
		name string
	}

	Wallet struct {
		balance Rupiah
	}
)

/**
Pointer in Function.
Each parameter in function, by default it's pass by value.
It's mean the data will be copied then send to function,
therefore, if we update a data in function, the original data will never change.
It's make the data is safe, but in some case, we need change the original data inside function,
to do this we can use pointer in function, to make parameter as pointer, we need "*" operator
*/

func updatePersonName(person *Person, name string) {
	person.name = name
}

/**
Pointer in Method.
Even though the method will be attached to the struct, the actual struct data accessed in the method is pass by value.
Very recommended to use pointer in the method, so it doesn't waste memory because it always has to be duplicated
when calling the method.
*/

func (w *Wallet) Deposit(amount Rupiah) {
	w.balance += amount
}

func (w *Wallet) Balance() Rupiah {
	return w.balance
}

/**
Nil.
Usually, If an object is not initialized, its value will automatically be null or nil.
It's different in Go-Lang, uninitialized variable with certain data type, its default value will automatically be created.
In Go-Lang, there is "nil", or in other words empty data.
Nil value can only be used in several data type: interface, function, map, slice, pointer, and channel.
*/

/**
Error Interface.
Go-Lang has interface to be used as contract for creating error, it's named "error".
To create error, we don't need to create it manually, Go-Lang has provided library to create helper easily.
*/

func (w *Wallet) Withdraw(amount Rupiah) error {
	if amount > w.balance {
		return errors.New("insufficient balance to make a withdrawal")
	}

	w.balance -= amount
	return nil
}

// Stringer, is a type that can describe itself as string
// For example, we can implement stringer on rupiah like this,
func (r Rupiah) String() string {
	return fmt.Sprintf("Rp.%d,-", r)
}

func main() {
	var person Person = Person{name: "Rosyid"}
	var person2 Person = person // it will copy person to person2

	person2.name = "Izzulkhaq"

	fmt.Printf("Person: %#v, address: %p \n", person, &person) // this data doesn't change
	fmt.Printf("Person2: %#v, address: %p \n", person2, &person2)

	fmt.Println()

	var person3 *Person = &person // it will reference to person
	person3.name = "M. Rosyid I"

	fmt.Printf("Person: %#v, address: %p \n", person, &person) // this data changed
	fmt.Printf("Person2: %#v, address: %p \n", person2, &person2)
	fmt.Printf("Person3: %#v, address: %p \n", person3, &person3)

	fmt.Println()

	*person3 = Person{name: "Rosyid Izz"}

	fmt.Printf("Person: %#v, address: %p \n", person, &person) // this data changed
	fmt.Printf("Person2: %#v, address: %p \n", person2, &person2)
	fmt.Printf("Person3: %#v, address: %p \n", person3, &person3)

	fmt.Println()

	// two ways to create pointer value
	var person4 *Person = &Person{name: "Rosyid"}
	var person5 *Person = new(Person)
	person5.name = "M. Rosyid"

	fmt.Printf("Person4: %#v, address: %p \n", person4, person4)
	fmt.Printf("Person5: %#v, address: %p \n", person5, person5)

	fmt.Println()

	updatePersonName(person5, "Rosyid I")
	updatePersonName(&person2, "Izz") // if variable isn't pointer, use "&" operator
	fmt.Printf("Person5: %#v, address: %p \n", person5, person5)
	fmt.Printf("Person2: %#v, address: %p \n", person2, &person2)
}
