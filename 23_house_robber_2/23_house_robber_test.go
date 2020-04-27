package houserobber

import "testing"

func TestHouseRobber(t *testing.T) {
	t.Run("Rob - should be 4", func(t *testing.T) {
		input := []int{1, 2, 3, 1}
		expected := 4
		actual := Rob(input)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})

	t.Run("Rob - should be 3", func(t *testing.T) {
		input := []int{2, 1, 1, 2}
		expected := 3
		actual := Rob(input)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})

	t.Run("Rob - should be 11", func(t *testing.T) {
		input := []int{2, 7, 9, 3, 1}
		expected := 11
		actual := Rob(input)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})

	t.Run("Rob - should be 14", func(t *testing.T) {
		input := []int{4, 1, 2, 7, 5, 3, 1}
		expected := 14
		actual := Rob(input)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})

	t.Run("Rob - should be 3", func(t *testing.T) {
		input := []int{1, 3, 1}
		expected := 3
		actual := Rob(input)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})

}
