package combinationsum

import "testing"

func TestLongestIncreasingSequence(t *testing.T) {
	t.Run("Combination sum - should be 4", func(t *testing.T) {
		input := []int{1, 2, 3}
		target := 4
		expected := 7
		actual := CombinationSum4(input, target)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})
	t.Run("Combination sum - should be 4", func(t *testing.T) {
		input := []int{4, 2, 1}
		target := 32
		expected := 39882198
		actual := CombinationSum4(input, target)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})

}
