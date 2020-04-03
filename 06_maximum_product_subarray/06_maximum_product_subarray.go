package maximumproductsubarray

//MaximumProductSubArray ...
func MaximumProductSubArray(array []int, left int, right int) int {
	if left == right {
		return array[left]
	}

	middle := (left + right) / 2

	l := MaximumProductSubArray(array, left, middle)
	r := MaximumProductSubArray(array, middle+1, right)
	c := crossMaximumProduct(array, left, right, middle)
	return max([]int{l, r, c})
}

func crossMaximumProduct(array []int, left int, right int, middle int) int {
	localProduct := 1
	maxLeftProduct := 0

	i := middle
	for i >= left {
		localProduct *= array[i]
		if localProduct > maxLeftProduct {
			maxLeftProduct = localProduct
		}

		i--
	}

	maxRightProduct := 0
	localProduct = 1
	i = middle + 1
	for i <= right {
		localProduct *= array[i]
		if localProduct > maxRightProduct {
			maxRightProduct = localProduct
		}

		i++
	}
	return maxLeftProduct * maxRightProduct
}

func max(array []int) int {
	max := 0

	i := 0
	for i < len(array) {
		if array[i] > max {
			max = array[i]
		}
		i++
	}

	return max
}
