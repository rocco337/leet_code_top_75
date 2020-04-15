package longestincreasingsequence

import "testing"

func TestLongestIncreasingSequence(t *testing.T) {
	t.Run("Longest increasing sequence should be 4", func(t *testing.T) {
		input := []int{10, 9, 2, 5, 3, 7, 101, 18}
		expected := 4
		actual := LengthOfLIS(input)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})
}
