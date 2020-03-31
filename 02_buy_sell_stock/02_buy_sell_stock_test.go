package buysellstock

import (
	"testing"
)

func TestBuySellStock(t *testing.T) {

	t.Run("Find max profit - should be 5", func(t *testing.T) {
		prices := []int{7, 1, 5, 3, 6, 4}
		expected := 5

		actual := MaxProfit(prices)
		if actual != expected {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}
	})

	t.Run("Find max profit - should be 0", func(t *testing.T) {
		prices := []int{77, 6, 4, 3, 1}
		expected := 0

		actual := MaxProfit(prices)
		if actual != expected {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}
	})
}

func BenchmarkBuySellStock(b *testing.B) {
	prices := []int{7, 1, 5, 3, 6, 4}
	expected := 5

	for n := 0; n < b.N; n++ {
		actual := MaxProfit(prices)
		if actual != expected {
			b.Error("Expected: ", expected, "Actual: ", actual)
		}
	}
}
