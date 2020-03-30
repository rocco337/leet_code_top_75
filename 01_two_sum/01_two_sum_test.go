package twosum

import (
	"testing"
)

func TestHuffman(t *testing.T) {

	t.Run("Find two numbers which sum equals to target", func(t *testing.T) {
		vector := []int{2, 7, 11, 15, 3}
		target := 9

		result := TwoSum(vector, target)

		if result[0]+result[1] != target {
			t.Error("Error", result)
		}
	})

	t.Run("Find two numbers which sum equals to target", func(t *testing.T) {
		vector := []int{0, 13, 2, 22, 3, 7, 11, 15}
		target := 9

		result := TwoSum(vector, target)

		if result[0]+result[1] != target {
			t.Error("Error", result)
		}
	})

	t.Run("Find two numbers which sum equals to target", func(t *testing.T) {
		vector := []int{0, 13, 7, 6, 22, 3, 2, 11, 15}
		target := 9

		result := TwoSum(vector, target)

		if result[0]+result[1] != target {
			t.Error("Error", result)
		}
	})
}

func BenchmarkTwoSum(b *testing.B) {
	vector := make([]int, 10000)
	i := 0
	for i < 10000 {
		vector[i] = i
		i++
	}
	target := 9543

	for n := 0; n < b.N; n++ {
		result := TwoSum(vector, target)
		if result[0]+result[1] != target {
			b.Error("Error", result)
		}
	}

}
