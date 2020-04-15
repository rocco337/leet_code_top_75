package hammingweight

import (
	"testing"
)

func TestHammingweight(t *testing.T) {
	t.Run("Get count of 1s in integer", func(t *testing.T) {
		expected := 3
		actual := HammingWeight(11)

		if actual != expected {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}
	})
}
