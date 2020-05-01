package decodeways

func NumDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	dp := make([]int, len(s))
	return numDecodings(s, 0, dp)
}

func numDecodings(s string, n int, dp []int) int {
	if n == len(s) {
		dp[len(s)-1] = 1
		return 1
	}

	if s[n] == '0' {
		return 0
	}

	if dp[n] != 0 {
		return dp[n]
	}

	count := 0
	if s[n] > '0' {
		count += numDecodings(s, n+1, dp)
	}

	if (n+2 <= len(s)) && ((s[n] == '2' && s[n+1] <= '6') || (s[n] == '1')) {
		count += numDecodings(s, n+2, dp)
	}

	return count
}
