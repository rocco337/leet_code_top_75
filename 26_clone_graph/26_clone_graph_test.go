package clonegraph

import "testing"

func TestUniquePaths(t *testing.T) {
	t.Run("Clone graph - should be equal", func(t *testing.T) {
		first := Node{}
		first.Val = 1

		second := Node{}
		second.Val = 2

		third := Node{}
		third.Val = 3

		forth := Node{}
		forth.Val = 4

		first.Neighbors = append(first.Neighbors, &second)
		first.Neighbors = append(first.Neighbors, &forth)

		second.Neighbors = append(second.Neighbors, &first)
		second.Neighbors = append(second.Neighbors, &third)

		third.Neighbors = append(third.Neighbors, &forth)
		third.Neighbors = append(third.Neighbors, &second)

		forth.Neighbors = append(forth.Neighbors, &first)
		forth.Neighbors = append(forth.Neighbors, &third)

		cloned := cloneGraph(&first)

		if cloned.Val != 1 || cloned.Neighbors[0].Val != 2 || cloned.Neighbors[1].Val != 4 {
			t.Error("Different in values in row: 0")
		}

	})

	t.Run("Clone 2 node graph - should be equal", func(t *testing.T) {
		first := Node{}
		first.Val = 1

		second := Node{}
		second.Val = 2

		first.Neighbors = append(first.Neighbors, &second)
		second.Neighbors = append(second.Neighbors, &first)

		cloned := cloneGraph(&first)

		if cloned.Val != 1 || cloned.Neighbors[0].Val != 2 {
			t.Error("Different in values in row: 0")
		}

		if second.Val != 2 || second.Neighbors[0].Val != 1 {
			t.Error("Different in values in row: 0")
		}

	})

	t.Run("Clone 1 node graph - should be equal", func(t *testing.T) {
		first := Node{}
		first.Val = 1

		cloned := cloneGraph(&first)

		if cloned.Val != 1 || len(cloned.Neighbors) != 0 {
			t.Error("Different in values in row: 0")
		}

	})

}
