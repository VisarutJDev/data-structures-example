package greedy

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	coins := []int{1, 5, 10, 25, 50}
	amount := 63
	result := CoinChange(coins, amount)
	if result != -1 {
		fmt.Printf("Minimum number of coins needed: %d\n", result)
	} else {
		fmt.Println("Change cannot be made with the given coins.")
	}
}
