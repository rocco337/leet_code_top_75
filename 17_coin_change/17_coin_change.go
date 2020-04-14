package coinchange

/*You are given coins of different denominations and a total amount of money amount.
Write a function to compute the fewest number of coins that you need to make up that amount.
If that amount of money cannot be made up by any combination of the coins, return -1.
*/
func CoinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	knapsack := make([]int, amount+1)

	for _, coin := range coins {
		i := 1
		for i <= amount {
			missing := i - coin

			if missing == 0 {
				knapsack[i] = 1
			} else if missing > 0 {
				hasComplement := knapsack[missing] != 0
				count := 1 + knapsack[missing]

				if (knapsack[i] == 0 && hasComplement) || (hasComplement && count < knapsack[i]) {
					knapsack[i] = count
				}
			}

			i++
		}

	}
	if knapsack[amount] == 0 {
		return -1
	}
	return knapsack[amount]
}
