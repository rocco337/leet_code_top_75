package countbits

import (
	"testing"
)

func TestSumOfTwoIntegers(t *testing.T) {
	t.Run("Count bits", func(t *testing.T) {
		expected := []int{0, 1, 1, 2, 1, 2}
		actual := CountBits(5)

		i := 0
		for i < len(actual) {
			if actual[i] != expected[i] {
				t.Error("Expected: ", expected, "Actual: ", actual)
				break
			}
			i++
		}

	})
}
