package climbingstairs

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	t.Run("ClimbStairs", func(t *testing.T) {
		n := 35
		expected := 14930352

		actual := ClimbStairs(n)

		if expected != actual {
			t.Error("Expected: ", expected, "Actual: ", actual)
		}

	})
}
