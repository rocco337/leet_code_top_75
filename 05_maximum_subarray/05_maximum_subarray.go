package maximumsubarray

//Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.
func MaximumSubArray(array []int) int {

	maxSum := array[0]
	maxSumUntilNow := array[0]

	i := 1
	for i < len(array) {

		if array[i] > maxSumUntilNow {
			maxSumUntilNow = array[i]
		} else {
			maxSumUntilNow += array[i]
		}

		if maxSumUntilNow > maxSum {
			maxSum = maxSumUntilNow
		}

		i++
	}
	return maxSum
}
