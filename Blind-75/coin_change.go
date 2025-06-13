package blind_75

/**
 * Author: Kelvin Le
 * Problem: 322. Coin Change
 * Link:https://leetcode.com/problems/coin-change/description/
 */

/**
 * Approach: Dynamic Programming
 * I realized that I can reach the amount by using the minimum number of coins at (amount - coin) + 1.
 * Therefore, I used a map to store the amount of coins need to reach of each from 0 to amount.
 * Time Complexity: O(n*amount)
 * Space Complexity: O(amount)
 */

func coinChange(coins []int, amount int) int {
	dp := make(map[int]int, amount+1)
	if amount == 0 {
		return 0
	}
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
