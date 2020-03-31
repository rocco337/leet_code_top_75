package containsduplicate

//ContainsDuplicate
//Given an array of integers, find if the array contains any duplicates.
//Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.
func ContainsDuplicate(array []int) bool {
	hashMap := make(map[int]int, 0)
	i := 0
	for i < len(array) {

		if _, contains := hashMap[array[i]]; contains {
			return true
		} else {
			hashMap[array[i]] = array[i]
		}
		i++
	}

	return false
}
