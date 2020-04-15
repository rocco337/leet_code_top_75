package longestincreasingsequence

//LengthOfLIS ...
//Given an unsorted array of integers, find the length of longest increasing subsequence.
func LengthOfLIS(nums []int) int {
	maxSequenceUntilNow := make([]int, len(nums))
	maxIncreasingSequence := 0

	i := 0
	for i < len(nums) {

		tempMax := 0
		j := 0
		for j < i {
			if nums[j] < nums[i] && maxSequenceUntilNow[j] > tempMax {
				tempMax = maxSequenceUntilNow[j]
			}
			j++
		}

		maxSequenceUntilNow[i] = tempMax + 1

		//set max incresing sequence so we dont need to iterate through array once more
		if maxSequenceUntilNow[i] > maxIncreasingSequence {
			maxIncreasingSequence = maxSequenceUntilNow[i]
		}
		i++
	}

	return maxIncreasingSequence
}
