package longestcommonsubsequence

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	t.Run("Longest common subsequence should be 3", func(t *testing.T) {
		input1 := "abcde"
		input2 := "ace"

		expected := 3
		actual := LongestCommonSubsequence(input1, input2)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})

	t.Run("Longest common subsequence should be 1", func(t *testing.T) {
		input1 := "psnw"
		input2 := "vozsh"

		expected := 1
		actual := LongestCommonSubsequence(input1, input2)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})

	t.Run("Longest common subsequence should be 2", func(t *testing.T) {
		input1 := "ezupkr"
		input2 := "ubmrapg"

		expected := 2
		actual := LongestCommonSubsequence(input1, input2)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})

	t.Run("Longest common subsequence should be 2", func(t *testing.T) {
		input1 := "oxcpqrsvwf"
		input2 := "shmtulqrypy"

		expected := 2
		actual := LongestCommonSubsequence(input1, input2)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})
}
