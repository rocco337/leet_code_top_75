package combinationsum

//CombinationSum4 ...
func CombinationSum4(nums []int, target int) int {
	if target == 0 || len(nums) == 0 {
		return 0
	}

	dp := make([]int, target+1)

	i := 1
	for i <= target {
		j := 0
		for j < len(nums) {
			if nums[j] == i {
				dp[i]++
			} else {
				complement := i - nums[j]
				if complement > 0 {
					dp[i] += dp[complement]
				}
			}

			j++
		}

		i++
	}
	return dp[target]
}
