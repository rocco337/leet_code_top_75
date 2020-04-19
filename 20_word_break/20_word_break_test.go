package wordbreak

import "testing"

func TestLongestIncreasingSequence(t *testing.T) {
	t.Run("Word break - should be true", func(t *testing.T) {
		expected := true
		actual := WordBreak("leetcode", []string{"leet", "code"})

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})

	t.Run("Word break - should be true", func(t *testing.T) {
		expected := true
		actual := WordBreak("applepenapple", []string{"apple", "pen"})

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})

	t.Run("Word break - should be false", func(t *testing.T) {
		expected := false
		actual := WordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})

	t.Run("Word break - should be true", func(t *testing.T) {
		expected := true
		actual := WordBreak("bb", []string{"a", "b", "bbb", "bbbb"})

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})

	t.Run("Word break - should be true", func(t *testing.T) {
		expected := true
		actual := WordBreak("aaaaaaa", []string{"aaaa", "aaa"})

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})
}
