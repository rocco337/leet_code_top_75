package productofarray

//ProductExceptSelf
//Given an array nums of n integers where n > 1,  return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].
func ProductExceptSelf(array []int) []int {
	result := make([]int, len(array))
	right := make([]int, len(array))

	result[0] = 1
	i := 1
	for i < len(array) {
		result[i] = array[i-1] * result[i-1]
		i++
	}

	i = len(array) - 2
	right[len(array)-1] = 1
	for i >= 0 {
		right[i] = array[i+1] * right[i+1]
		i--
	}

	i = 0
	for i < len(array) {
		result[i] = result[i] * right[i]
		i++
	}

	return result
}
