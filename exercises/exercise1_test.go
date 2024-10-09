package exercises

import (
	"testing"
)

func TestExercies(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7}
	expected := []int{2, 4, 6}
	t.Run("test Even function", func(t *testing.T) {

		expectedLength := len(expected)
		result := EvenNumbers(list)
		actualLength := len(result)
		t.Log(result)
		var valid bool = true

		if expectedLength != actualLength {
			t.Error("result has incorrect number of items")
		}

		for i := range expected {
			if expected[i] != result[i] {
				valid = false
			}
		}
		if !valid {
			t.Error("result has incorrect items other than expected")
		}
	})

	t.Run("test Odd function", func(t *testing.T) {
		expected := []int{1, 3, 5, 7}
		expectedLength := len(expected)
		result := OddNumbers(list)
		actualLength := len(result)
		t.Log(result)
		var valid bool = true

		if expectedLength != actualLength {
			t.Error("result has incorrect number of items")
		}

		for i := range expected {
			if expected[i] != result[i] {
				valid = false
			}
		}
		if !valid {
			t.Error("result has incorrect items other than expected")
		}
	})

	t.Run("test Primer function", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
		expected := []int{2, 3, 5, 7, 11}
		expectedLength := len(expected)
		result := PrimeNumbers(list)
		actualLength := len(result)
		t.Log(result)
		var valid bool = true

		if expectedLength != actualLength {
			t.Error("result has incorrect number of items")
		}

		for i := range expected {
			if expected[i] != result[i] {
				valid = false
			}
		}
		if !valid {
			t.Error("result has incorrect items other than expected")
		}
	})

	t.Run("test Prime which is ODD function", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
		expected := []int{3, 5, 7, 11}
		expectedLength := len(expected)
		result := PrimeNumbersAndOdd(list)
		actualLength := len(result)
		t.Log(result)
		var valid bool = true

		if expectedLength != actualLength {
			t.Error("result has incorrect number of items")
		}

		for i := range expected {
			if expected[i] != result[i] {
				valid = false
			}
		}
		if !valid {
			t.Error("result has incorrect items other than expected")
		}
	})

	t.Run("test Even and Mulitples of 5", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 20, 25, 46}
		expected := []int{10, 20}
		expectedLength := len(expected)
		result := EvenAndMulitplesofFive(list)
		actualLength := len(result)
		t.Log(result)
		var valid bool = true

		if expectedLength != actualLength {
			t.Error("result has incorrect number of items")
		}

		for i := range expected {
			if expected[i] != result[i] {
				valid = false
			}
		}
		if !valid {
			t.Error("result has incorrect items other than expected")
		}
	})

	t.Run("test match ALL the conditions.", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 20, 25, 46}
		expected := []int{10, 20}
		expectedLength := len(expected)
		result1 := FunctionHandler(list, OddNumbers)
		actualLength := len(result1)
		t.Log(result1)
		var valid bool = true

		if expectedLength != actualLength {
			t.Error("result has incorrect number of items")
		}

		for i := range expected {
			if expected[i] != result[i] {
				valid = false
			}
		}
		if !valid {
			t.Error("result has incorrect items other than expected")
		}
	})

}
