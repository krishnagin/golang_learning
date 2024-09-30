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

}
