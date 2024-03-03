package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying Hello World", func(t *testing.T) {
		got := Hello()
		want := "Hello World!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
