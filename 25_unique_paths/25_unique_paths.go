package uniquepaths

func UniquePaths(m int, n int) int {
	dp := make([][]int, n+1)

	i := 1
	for i <= n {
		dp[i] = make([]int, m+1)
		i++
	}
	return uniquePathsRecursive(m, n, 1, 1, dp)
}

func uniquePathsRecursive(m int, n int, x int, y int, dp [][]int) int {
	if x == m && y == n {
		dp[y][x] = 1
		return 1
	}
	if dp[y][x] != 0 {
		return dp[y][x]
	}

	count := 0
	if x+1 <= m {
		count += uniquePathsRecursive(m, n, x+1, y, dp)
		dp[y][x] = count
	}

	if y+1 <= n {
		count += uniquePathsRecursive(m, n, x, y+1, dp)
		dp[y][x] = count
	}
	return count
}
