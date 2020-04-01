package productofarray

import "testing"

func TestProductOfArray(t *testing.T) {
	t.Run("ProductOfArray - should return", func(t *testing.T) {
		input := []int{1, 2, 3, 4}
		expected := []int{24, 12, 8, 6}

		actual := ProductExceptSelf(input)

		if len(actual) != len(expected) {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}

		i := 0
		for i < len(actual) {
			if actual[i] != expected[i] {
				t.Error("Expected: ", expected, "Actual: ", actual)
			}
			i++
		}

	})
}
