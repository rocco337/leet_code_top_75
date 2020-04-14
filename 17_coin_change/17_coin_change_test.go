package coinchange

import "testing"

func TestCoinChange(t *testing.T) {
	t.Run("Coin change -min should be 3", func(t *testing.T) {
		input := []int{1, 2, 5}
		target := 11
		expected := 3
		actual := CoinChange(input, target)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})
}
