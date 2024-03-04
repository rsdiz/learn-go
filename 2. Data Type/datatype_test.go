package main

import (
	"math/rand"
	"reflect"
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
		numbers := make([]int, rand.Intn(15))

		want := 0
		for index, _ := range numbers {
			numbers[index] = rand.Intn(100)
			want += numbers[index]
		}
		t.Logf("Numbers: %v", numbers)
		t.Logf("Expected: %d", want)

		got := Sum(numbers)

		assertCorrectIntValue(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Sum Multiple Data", func(t *testing.T) {
		got := SumAll(
			[]int{4, 7, 9},
			[]int{5, 4, 1},
		)
		expected := []int{20, 10}
		t.Logf("Got: %v", got)
		t.Logf("Expected: %d", expected)

		assertCorrectSliceValue(t, got, expected)
	})

	t.Run("Sum Random Multiple Data", func(t *testing.T) {
		data1 := make([]int, rand.Intn(15))
		data2 := make([]int, rand.Intn(15))
		expected := make([]int, 2)

		want := 0
		for index, _ := range data1 {
			data1[index] = rand.Intn(100)
			want += data1[index]
		}
		expected[0] = want

		want = 0
		for index, _ := range data2 {
			data2[index] = rand.Intn(100)
			want += data2[index]
		}
		expected[1] = want

		got := SumAll(data1, data2)
		t.Logf("Data 1: %v", data1)
		t.Logf("Data 2: %v", data2)
		t.Logf("Got: %v", got)
		t.Logf("Expected: %d", expected)

		assertCorrectSliceValue(t, got, expected)
	})
}

func assertCorrectIntValue(t *testing.T, sum, expected int) {
	t.Helper()
	if !reflect.DeepEqual(sum, expected) {
		t.Errorf("sum %d expected %d", sum, expected)
	}
}

func assertCorrectFloatValue(t *testing.T, sum, expected float64) {
	t.Helper()
	if !reflect.DeepEqual(sum, expected) {
		t.Errorf("sum %f expected %f", sum, expected)
	}
}

func assertCorrectSliceValue(t *testing.T, got, expected []int) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v expected %v", got, expected)
	}
}
