package sumoftwointegers

import (
	"testing"
)

func TestSumOfTwoIntegers(t *testing.T) {
	t.Run("Add two numbers", func(t *testing.T) {
		expected := 5
		actual := GetSum(2, 3)

		if actual != expected {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}
	})
}
