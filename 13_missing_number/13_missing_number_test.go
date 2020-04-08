package missingnumber

import (
	"testing"
)

func TestMissingNumber(t *testing.T) {
	t.Run("Should find missing number", func(t *testing.T) {
		input := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
		expected := 8
		actual := MissingNumber(input)

		if actual != expected {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}
	})

	t.Run("Should find missing number", func(t *testing.T) {
		input := []int{0}
		expected := 1
		actual := MissingNumber(input)

		if actual != expected {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}
	})

	t.Run("Should find missing number", func(t *testing.T) {
		input := []int{1}
		expected := 0
		actual := MissingNumber(input)

		if actual != expected {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}
	})
}
