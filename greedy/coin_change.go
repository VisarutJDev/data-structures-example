package greedy

import (
	"math"
)

// CoinChange calculates the minimum number of coins needed to make up the amount
func CoinChange(coins []int, amount int) int {
	// Create a table to store the minimum number of coins for each amount
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0 // Base case: 0 amount requires 0 coins

	// Update the table with minimum coins required for each amount
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if dp[i-coin] != math.MaxInt {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] == math.MaxInt {
		return -1 // If amount cannot be made with given coins
	}
	return dp[amount]
}

// min returns the smaller of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
