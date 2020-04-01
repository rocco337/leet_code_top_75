package maximumsubarray

import "testing"

func TestMaxiumumSubArray(t *testing.T) {

	t.Run("Find max sub array", func(t *testing.T) {
		input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
		expected := 6

		actual := MaximumSubArray(input)
		if actual != expected {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}
	})
}
