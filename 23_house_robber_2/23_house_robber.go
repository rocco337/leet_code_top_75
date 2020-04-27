package houserobber

import "fmt"

func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	result := max(robInternal(nums[0:len(nums)-1]), robInternal(nums[1:]))

	return result
}

func robInternal(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	i := 0
	for i < len(nums) {
		fmt.Println("i", i, "len", len(nums))
		dp[i] = max(dp[i], nums[i])

		if i+2 >= len(nums) {
			break
		}

		if i+2 < len(nums) {
			dp[i+2] = max(dp[i+2], dp[i]+nums[i+2])
		}

		if i-1 >= 0 {
			dp[i+2] = max(dp[i+2], dp[i-1]+nums[i+2])
		}

		i++
	}

	return max(dp[len(nums)-2], dp[len(nums)-1])
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
