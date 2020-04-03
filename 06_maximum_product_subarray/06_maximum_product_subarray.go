package maximumproductsubarray

//MaximumProductSubArray ...
func MaximumProductSubArray(array []int) int {

	if len(array) == 1 {
		return array[0]
	}

	globalMax := 0
	maxUntilNow := 0
	minUntilNow := 0
	i := 0
	for i < len(array) {
		temp := maxUntilNow

		maxUntilNow = max(array[i], max(array[i]*maxUntilNow, array[i]*minUntilNow))
		minUntilNow = min(array[i], min(array[i]*temp, array[i]*minUntilNow))

		globalMax = max(maxUntilNow, globalMax)

		i++
	}
	return globalMax
}

func max(left int, right int) int {
	if left > right {
		return left
	}

	return right
}

func min(left int, right int) int {
	if left < right {
		return left
	}

	return right
}
