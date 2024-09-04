# Greedy algorithm
Greedy algorithms work by making a sequence of choices, each of which looks best at the moment, with the hope that these choices will lead to an optimal solution. They are typically used in optimization problems.

### Coin change problem
The goal is to find the minimum number of coins needed to make a given amount of change. This example assumes the coin denominations are in a set where a greedy approach is optimal (e.g., denominations like 1, 5, 10, 25, 50).

**Explanation:**
The dp array stores the minimum number of coins needed for each amount up to the target amount.
For each coin and each possible amount, the algorithm updates the dp array to find the minimum number of coins required.

### Activity Selection Problem
The goal is to select the maximum number of activities that don't overlap, given their start and end times. This problem is solved efficiently using a greedy approach by always selecting the next activity that ends the earliest.

**Explanation:**
The ActivitySelection function sorts the activities by their end times.
It iterates through the sorted list and selects an activity if it starts after the last selected activity ends.