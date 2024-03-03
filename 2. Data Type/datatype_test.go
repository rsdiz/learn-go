package main

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(2, 3)
	expected := 2 + 3

	if sum != expected {
		t.Errorf("sum %d expected %d", sum, expected)
	}
}

func TestReduction(t *testing.T) {
	sum := Reduction(211, 74)
	expected := 211 - 74

	if sum != expected {
		t.Errorf("sum %d expected %d", sum, expected)
	}
}

func TestMultiply(t *testing.T) {
	sum := Multiply(32.0, 21.0)
	expected := 32.0 * 21.0

	if sum != expected {
		t.Errorf("sum %f expected %f", sum, expected)
	}
}

func TestDivider(t *testing.T) {
	sum := Divider(318, 24.0)
	expected := 318 / 24.0

	if sum != expected {
		t.Errorf("sum %f expected %f", sum, expected)
	}
}
