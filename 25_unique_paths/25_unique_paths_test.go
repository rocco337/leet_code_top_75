package uniquepaths

import "testing"

func TestUniquePaths(t *testing.T) {
	t.Run("Unique paths - should be 1", func(t *testing.T) {
		expected := 1
		actual := UniquePaths(1, 1)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})

	t.Run("Unique paths - should be 3", func(t *testing.T) {
		expected := 3
		actual := UniquePaths(3, 2)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})

	t.Run("Unique paths - should be 28", func(t *testing.T) {
		expected := 28
		actual := UniquePaths(7, 3)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})

}
