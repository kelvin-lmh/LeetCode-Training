package blind_75

/**
 * Author: Kelvin Le
 * Problem: 198. House Robber
 * Link:https://leetcode.com/problems/house-robber/description/
 */

/**
 * Approach: Dynamic Programming
 * I use DP to keep track of the maximum amount of money that can be robbed up to each house.
 * I remain two variable (prev1 , prev2) to store the maximum amount robbed from the previous two houses
 * Time Complexity: O(n)
 * Space Complexity: O(1)
 */

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	cur, prev1, prev2 := 0, 0, 0
	for _, num := range nums {
		cur = max(prev1, prev2+num)
		prev2, prev1 = prev1, cur
	}
	return prev1
}
