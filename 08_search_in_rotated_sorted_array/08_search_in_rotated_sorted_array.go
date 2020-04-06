package searchinrotatedsortedarray

/*
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).
*/
func Search(array []int, target int) int {
	if len(array) == 0 {
		return -1
	}

	pivotIndex := findPivotIndex(array, 0, len(array)-1)

	left := 0
	right := len(array) - 1
	if pivotIndex != -1 {
		if target <= array[pivotIndex-1] && target >= array[0] {
			right = pivotIndex - 1
		} else {
			left = pivotIndex
		}
	}

	return binarySearch(array, left, right, target)
}

func binarySearch(array []int, left int, right int, target int) int {
	if left > right {
		return -1
	}

	if left == right {
		if array[left] == target {
			return left
		}

		return -1
	}

	middle := (left + right) / 2

	if array[middle] == target {
		return middle
	}

	if target > array[middle] {
		return binarySearch(array, middle+1, right, target)
	}

	return binarySearch(array, left, middle-1, target)

}

func findPivotIndex(array []int, left int, right int) int {
	if left == right {
		return -1
	}

	middle := (left + right) / 2
	if array[middle] > array[middle+1] {
		return middle + 1
	}

	l := findPivotIndex(array, left, middle)
	if l != -1 {
		return l
	}

	return findPivotIndex(array, middle+1, right)
}
