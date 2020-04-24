package houserobber

func Rob(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	i := 0
	for i < len(nums)-1 {
		if nums[i] > dp[i] {
			dp[i] = nums[i]
		}

		if i+2 >= len(nums) {
			break
		}
		next := nums[i+2]

		if i+2 < len(nums) && dp[i]+nums[i+2] > dp[i+2] {
			dp[i+2] = dp[i] + nums[i+2]
		}

		if i-1 >= 0 && dp[i-1]+next > dp[i+2] {
			dp[i+2] = dp[i-1] + nums[i+2]
		}

		i++
	}

	if dp[len(nums)-2] > dp[len(nums)-1] {
		return dp[len(nums)-2]
	}

	return dp[len(nums)-1]

}
