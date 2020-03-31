package containsduplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	t.Run("Contains duplicate - should return true", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 4, 7, 8, 9}
		expected := true

		actual := ContainsDuplicate(input)

		if actual != expected {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}
	})

	t.Run("Contains duplicate - should return false", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		expected := false

		actual := ContainsDuplicate(input)

		if actual != expected {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}
	})
}

func BenchmarkContainsDuplicate(b *testing.B) {
	input := []int{1, 2, 3, 4, 5, 6, 4, 7, 8, 9}
	expected := true

	for n := 0; n < b.N; n++ {
		actual := ContainsDuplicate(input)

		if actual != expected {
			b.Error("Expected: ", expected, "Actual: ", actual)
		}
	}
}
