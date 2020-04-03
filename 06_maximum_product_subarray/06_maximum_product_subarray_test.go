package maximumproductsubarray

import (
	"testing"
)

func TestMaximumProductSubarray(t *testing.T) {
	t.Run("Array of two negative numbers - should give product of positive number", func(t *testing.T) {
		input := []int{-4, -3}
		expected := 12

		actual := MaximumProductSubArray(input)

		if actual != expected {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}
	})

	t.Run("Maximum product should be as expected", func(t *testing.T) {

		input := []int{2, 3, -2, 4}
		expected := 6

		actual := MaximumProductSubArray(input)

		if actual != expected {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}
	})

	t.Run("Maximum product should be as expected", func(t *testing.T) {

		input := []int{-2, 3, -4}
		expected := 24

		actual := MaximumProductSubArray(input)

		if actual != expected {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}
	})
}
