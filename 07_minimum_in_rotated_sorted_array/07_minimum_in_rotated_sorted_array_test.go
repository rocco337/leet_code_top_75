package minimuminrotatedsortedarray

import "testing"

func TestFindMinimumInRotatedSortedArray(t *testing.T) {
	t.Run("Find minimum in rotated sorted array", func(t *testing.T) {
		input := []int{4, 5, 6, 7, 0, 1, 2}
		expected := 0

		actual := FindMinumum(input)
		if actual != expected {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}
	})
}

func BenchmarkFindMinumum(b *testing.B) {
	input := []int{3, 1, 2}
	expected := 3

	for n := 0; n < b.N; n++ {
		actual := FindMinumum(input)
		if actual != expected {
			b.Error("Expected: ", expected, "Actual: ", actual)
		}
	}
}
