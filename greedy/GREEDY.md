# Greedy algorithm

### Coin change problem
The goal is to find the minimum number of coins needed to make a given amount of change. This example assumes the coin denominations are in a set where a greedy approach is optimal (e.g., denominations like 1, 5, 10, 25, 50).

**Explanation:**
The dp array stores the minimum number of coins needed for each amount up to the target amount.
For each coin and each possible amount, the algorithm updates the dp array to find the minimum number of coins required.