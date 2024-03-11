package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Rupiah) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Rupiah(1000))
		assertBalance(t, wallet, Rupiah(1000))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Rupiah(1000)}
		_ = wallet.Withdraw(Rupiah(500))
		assertBalance(t, wallet, Rupiah(500))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Rupiah(500)}
		err := wallet.Withdraw(Rupiah(750))
		assertError(t, err)
		assertBalance(t, wallet, Rupiah(500))
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
