package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	perimeterTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, want: 40.0},
		{name: "circle", shape: Circle{Radius: 10.0}, want: 62.83185307179586},
	}

	for _, tt := range perimeterTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			assertCorrectFloatValue(t, tt.shape, got, tt.want)
		})
	}
}

func TestArea(t *testing.T) {
	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, want: 100.0},
		{name: "circle", shape: Circle{Radius: 10.0}, want: 314.1592653589793},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			assertCorrectFloatValue(t, tt.shape, got, tt.want)
		})
	}
}

func TestUser_SayHi(t *testing.T) {
	user := User{
		fullName: "Rosyid",
	}
	user2 := User{
		fullName: "Izzulkhaq",
	}
	got := user.SayHi(user2)
	want := "Hi Izzulkhaq, my name is Rosyid"

	assertCorrectStringValue(t, got, want)
}

func assertCorrectFloatValue(t testing.TB, shape Shape, got, want float64) {
	t.Helper()

	if got != want {
		t.Errorf("%#v got %g want %g", shape, got, want)
	}
}

func assertCorrectStringValue(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
