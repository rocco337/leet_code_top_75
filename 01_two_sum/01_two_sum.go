package twosum

//TwoSum ...
//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
func TwoSum(vector []int, target int) []int {
	if len(vector) < 2 {
		return []int{}
	}
	hashMap := make(map[int]int, 0)
	i := 0
	for i < len(vector) {
		complement := target - vector[i]
		if _, found := hashMap[complement]; found {
			return []int{vector[i], complement}
		}
		hashMap[vector[i]] = vector[i]
		i++
	}

	return make([]int, 0)
}
