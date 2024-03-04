package main

import (
	"math/rand"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("Simple Adding Test", func(t *testing.T) {
		sum := Add(2, 3)
		expected := 2 + 3

		assertCorrectIntValue(t, sum, expected)
	})

	t.Run("More Adding Test", func(t *testing.T) {
		sum := Add(4124, 9475)
		expected := 4124 + 9475

		assertCorrectIntValue(t, sum, expected)
	})
}

func TestReduction(t *testing.T) {
	t.Run("Simple Reduction Test", func(t *testing.T) {
		sum := Reduction(211, 74)
		expected := 211 - 74

		assertCorrectIntValue(t, sum, expected)
	})

	t.Run("More Reduction Test", func(t *testing.T) {
		sum := Reduction(92345, 3642)
		expected := 92345 - 3642

		assertCorrectIntValue(t, sum, expected)
	})
}

func TestMultiply(t *testing.T) {
	t.Run("Simple Multiply Test", func(t *testing.T) {
		sum := Multiply(32.0, 21.0)
		expected := 32.0 * 21.0

		assertCorrectFloatValue(t, sum, expected)
	})

	t.Run("More Multiply Test", func(t *testing.T) {
		sum := Multiply(94.8, 74)
		expected := 94.8 * 74

		assertCorrectFloatValue(t, sum, expected)
	})
}

func TestDivider(t *testing.T) {
	t.Run("Simple Divider Test", func(t *testing.T) {
		sum := Divider(318, 24.0)
		expected := 318 / 24.0

		assertCorrectFloatValue(t, sum, expected)
	})

	t.Run("More Divider Test", func(t *testing.T) {
		sum := Divider(9026, 21.5)
		expected := 9026 / 21.5

		assertCorrectFloatValue(t, sum, expected)
	})
}

func TestSum(t *testing.T) {
	t.Run("Sum number 1 until 5", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers[:])
		want := 15

		assertCorrectIntValue(t, got, want)
	})

	t.Run("Sum random numbers", func(t *testing.T) {
		numbers := make([]int, 10)

		want := 0
		for index, _ := range numbers {
			numbers[index] = rand.Intn(100)
			want += numbers[index]
		}
		t.Log("Numbers:", numbers)
		t.Log("Expected:", want)

		got := Sum(numbers)

		assertCorrectIntValue(t, got, want)
	})
}

func assertCorrectIntValue(t *testing.T, sum, expected int) {
	t.Helper()
	if sum != expected {
		t.Errorf("sum %d expected %d", sum, expected)
	}
}

func assertCorrectFloatValue(t *testing.T, sum float64, expected float64) {
	t.Helper()
	if sum != expected {
		t.Errorf("sum %f expected %f", sum, expected)
	}
}
