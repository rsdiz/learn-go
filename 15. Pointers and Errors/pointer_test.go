package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Rupiah(1000))
		assertBalance(t, wallet, Rupiah(1000))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Rupiah(1000)}
		err := wallet.Withdraw(Rupiah(500))
		assertNoError(t, err)
		assertBalance(t, wallet, Rupiah(500))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Rupiah(500)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Rupiah(750))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
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

func assertBalance(t testing.TB, wallet Wallet, want Rupiah) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatalf("wanted an error but didn't get one")
	}

	if !errors.Is(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}
