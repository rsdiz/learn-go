package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Rupiah(1000))

	got := wallet.Balance()
	fmt.Printf("address of balance in test is %p \n", &wallet.balance)

	want := Rupiah(1000)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPerson(t *testing.T) {
	person := Person{name: "Izzulkhaq"}
	updatePersonName(&person, "Rosyid")

	got := person.name
	fmt.Printf("address of person in test is %p \n", &person.name)
	want := "Rosyid"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
