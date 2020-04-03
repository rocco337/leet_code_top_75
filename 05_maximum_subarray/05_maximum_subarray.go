package maximumsubarray

import (
	"math"
)

//Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
func MaximumSubArray(array []int, left int, right int) int {

	if left == right {
		return array[left]
	}

	middlePoint := (left + right) / 2

	l := MaximumSubArray(array, left, middlePoint)
	r := MaximumSubArray(array, middlePoint+1, right)
	c := maxCrossingSubArray(array, left, middlePoint, right)
	return max([]int{l, r, c})
}

func maxCrossingSubArray(array []int, left int, mid int, right int) int {

	leftSum := math.MinInt32
	sum := 0
	i := mid
	for i >= left {
		sum += array[i]
		if sum > leftSum {
			leftSum = sum
		}

		i--
	}

	i = mid + 1
	rightSum := math.MinInt32
	sum = 0
	for i <= right {
		sum += array[i]
		if sum > rightSum {
			rightSum = sum
		}
		i++
	}

	return leftSum + rightSum

}

func max(array []int) int {
	max := math.MinInt32

	i := 0
	for i < len(array) {
		if array[i] > max {
			max = array[i]
		}
		i++
	}

	return max
}
