package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Rupiah(1000))

		got := wallet.Balance()

		want := Rupiah(1000)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Rupiah(1000)}
		wallet.Withdraw(Rupiah(500))

		got := wallet.Balance()

		want := Rupiah(500)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
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
