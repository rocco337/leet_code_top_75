package wordbreak

func WordBreak(s string, wordDict []string) bool {
	if len(wordDict) == 0 {
		return false
	}

	if len(wordDict) == 0 {
		return false
	}

	dict := make(map[string]int, len(wordDict))
	for _, word := range wordDict {
		dict[word] = len(word)
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	i := 1
	for i <= len(s) {
		j := i - 1
		for j >= 0 {
			_, exist := dict[s[j:i]]

			if dp[j] && exist {
				dp[i] = true
				break
			}

			j--
		}
		i++
	}
	return dp[len(s)]
}
