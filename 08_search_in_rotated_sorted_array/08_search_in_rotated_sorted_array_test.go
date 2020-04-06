package searchinrotatedsortedarray

import "testing"

func TestSearchInSortedArray(t *testing.T) {
	t.Run("Should find target on index", func(t *testing.T) {
		input := []int{4, 5, 6, 7, 0, 1, 2}
		target := 0
		expected := 4
		actual := Search(input, target)

		if actual != expected {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}
	})

	t.Run("Target doenst exist, should return -1", func(t *testing.T) {
		input := []int{4, 5, 6, 7, 0, 1, 2}
		target := 10
		expected := -1
		actual := Search(input, target)

		if actual != expected {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}
	})

	t.Run("Only one item in array, and it matches target", func(t *testing.T) {
		input := []int{1}
		target := 1
		expected := 0
		actual := Search(input, target)

		if actual != expected {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}
	})

	t.Run("Two items in array, should find target", func(t *testing.T) {
		input := []int{3, 1}
		target := 3
		expected := 0
		actual := Search(input, target)

		if actual != expected {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}
	})
}
