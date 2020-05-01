package decodeways

import "testing"

func TestDecodeWays(t *testing.T) {
	t.Run("NumDecodings - should be 3", func(t *testing.T) {
		input := "226"
		expected := 3
		actual := NumDecodings(input)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})

	t.Run("NumDecodings - should be 2", func(t *testing.T) {
		input := "17"
		expected := 2
		actual := NumDecodings(input)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})

	t.Run("NumDecodings - should be 5", func(t *testing.T) {
		input := "32119"
		expected := 5
		actual := NumDecodings(input)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})

	t.Run("NumDecodings - should be 1", func(t *testing.T) {
		input := "110"
		expected := 1
		actual := NumDecodings(input)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})
}
