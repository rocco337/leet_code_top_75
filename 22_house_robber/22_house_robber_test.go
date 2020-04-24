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

	t.Run("Rob - should be 4", func(t *testing.T) {
		input := []int{2, 1, 1, 2}
		expected := 4
		actual := Rob(input)

		if actual != expected {
			t.Error("Expected:", expected, "Actual: ", actual)
		}
	})

	t.Run("Rob - should be 12", func(t *testing.T) {
		input := []int{2, 7, 9, 3, 1}
		expected := 12
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

	//[226,174,214,16,218,48,153,131,128,17,157,142,88,43,37,157,43,221,191,68,206,23,225,82,54,118,111,46,80,49,245,63,25,194,72,80,143,55,209,18,55,122,65,66,177,101,63,201,172,130,103,225,142,46,86,185,62,138,212,192,125,77,223,188,99,228,90,25,193,211,84,239,119,234,85,83,123,120,131,203,219,10,82,35,120,180,249,106,37,169,225,54,103,55,166,124]

}
