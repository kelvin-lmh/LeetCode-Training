package blind_75

/**
 * Author: Kelvin Le
 * Problem: 213. House Robber II
 * Link:https://leetcode.com/problems/house-robber-ii/description/
 */

/**
 * Approach: Dynamic Programming
 * I use a dynamic programming approach similar to House Robber I, but I handle the circular nature of the houses.
 * I calculate the maximum amount that can be robbed in two scenarios:
 * 1. Robbing from the first house to the second last house.
 * 2. Robbing from the second house to the last house.
 * The final result is the maximum of these two scenarios.
 * Time Complexity: O(n)
 * Space Complexity: O(1)
 */

func rob2(nums []int) int {
	n := len(nums)

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	robLinear := func(numbers []int) int {
		cur, prev1, prev2 := 0, 0, 0
		for _, num := range numbers {
			cur = max(prev1, prev2+num)
			prev1, prev2 = cur, prev1
		}
		return prev1
	}
	case1 := robLinear(nums[:n-1])
	case2 := robLinear(nums[1:n])

	return max(case1, case2)
}
