package minimuminrotatedsortedarray

//uppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//(i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
//Find the minimum element.
//You may assume no duplicate exists in the array.
func FindMinumum(array []int) int {

	if array[0] < array[len(array)-1] {
		//its just sorted array withou pivoting
		return array[0]
	}

	return FindMinumumRecursive(array, 0, len(array)-1)
}

func FindMinumumRecursive(array []int, left int, right int) int {
	if left == right {
		return array[left]
	}

	middle := (left + right) / 2

	if array[middle] > array[middle+1] {
		return array[middle+1]
	}

	l := FindMinumumRecursive(array, left, middle)
	r := FindMinumumRecursive(array, middle+1, right)

	return min(l, r)
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
