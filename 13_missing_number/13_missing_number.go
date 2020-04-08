package missingnumber

//Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.
func MissingNumber(nums []int) int {
	hash := make(map[int]int, len(nums))
	i := 0
	for i < len(nums) {
		hash[nums[i]] = i
		i++
	}

	i = 0
	for i < len(nums)+1 {
		if _, exists := hash[i]; !exists {
			return i
		}
		i++
	}

	panic("not found")

}
