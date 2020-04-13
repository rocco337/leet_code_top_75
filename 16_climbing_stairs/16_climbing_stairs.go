package climbingstairs

/* You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.
*/
func ClimbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	result := make([]int, n+1)
	result[1] = 1
	result[2] = 2

	i := 3
	for i <= n {
		result[i] = result[i-2] + result[i-1]
		i++
	}

	return result[n]
}
