package blind_75

/**
 * Author: Kelvin Le
 * Problem: 377. Combination Sum IV
 * Link:https://leetcode.com/problems/combination-sum-iv/description/
 */

/**
 * Approach: Dynamic Programming
 * I use a dynamic programming approach to count the number of combinations that sum up to the target.
 * I maintain a dp array where dp[i] represents the number of combinations that sum up to i.
 * Time Complexity: O(n * target), where n is the length of the input array.
 * Space Complexity: O(target), for the dp array.
 */

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
