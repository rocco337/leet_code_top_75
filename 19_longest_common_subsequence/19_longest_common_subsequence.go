package longestcommonsubsequence

//LongestCommonSubsequence ...
func LongestCommonSubsequence(text1 string, text2 string) int {

	result := make([][]int, len(text2)+1)

	//initialize 0 row, so we can take easily 0 value as previous
	result[0] = make([]int, len(text1)+1)

	row := 1
	for row <= len(text2) {
		result[row] = make([]int, len(text1)+1)

		column := 1
		for column <= len(text1) {
			if text1[column-1] == text2[row-1] {
				result[row][column] = result[row-1][column-1] + 1
			} else {
				result[row][column] = max(result[row-1][column-1], result[row-1][column], result[row][column-1])
			}
			column++
		}
		row++
	}

	return result[len(text2)][len(text1)]
}

func max(x int, y int, z int) int {
	if x > y && x > z {
		return x
	}

	if y > z {
		return y
	}

	return z
}
