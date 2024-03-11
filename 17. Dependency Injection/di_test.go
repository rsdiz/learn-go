package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Rosyid")

	got := buffer.String()
	want := "Hi, Rosyid"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
