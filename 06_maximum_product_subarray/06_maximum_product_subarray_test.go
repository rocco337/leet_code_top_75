package maximumproductsubarray

import "testing"

func TestMaximumProductSubarray(t *testing.T) {
	t.Run("Maximum product should be as expected", func(t *testing.T) {
		input := []int{-1, 2, 3, -2, 4}
		expected := 6

		actual := MaximumProductSubArray(input, 0, len(input)-1)

		if actual != expected {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}
	})
}
