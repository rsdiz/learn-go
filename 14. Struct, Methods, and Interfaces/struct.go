package main

import (
	"fmt"
	"time"
)

/**
Struct is a data template used to combine zero or more other data types into one unit.
Struct usually the representation of data in the programs we create.
Every data in structs are saved in the field.
We can call struct is collection of fields.
*/

type (
	User struct {
		fullName    string
		dateOfBirth time.Time
		phoneNumber string
	}

	Rectangle struct {
		Width  float64
		Height float64
	}
)

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

func main() {
	// First way to initialize struct
	location, _ := time.LoadLocation("Asia/Jakarta")
	user := User{
		fullName:    "Rosyid",
		dateOfBirth: time.Date(2000, 5, 24, 0, 0, 0, 0, location),
		phoneNumber: "08123456789",
	}
	fmt.Printf("Name: %s\nDate of Birth: %s\nPhone Number: %s\n", user.fullName, user.dateOfBirth.Format("Monday, 2 January 2006"), user.phoneNumber)

	// Second way to initialize struct
	user2 := User{}
	user2.fullName = "Izzulkhaq"
	user2.dateOfBirth = time.Date(2008, 7, 1, 0, 0, 0, 0, location)
	user2.phoneNumber = "08987654321"

	fmt.Printf("Name: %s\nDate of Birth: %s\nPhone Number: %s\n", user2.fullName, user2.dateOfBirth.Format("Monday, 2 January 2006"), user2.phoneNumber)
}
