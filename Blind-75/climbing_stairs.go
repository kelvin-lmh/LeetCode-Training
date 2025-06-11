package blind_75

/**
 * Author: Kelvin Le
 * Problem: 70. Climbing Stairs
 * Link:https://leetcode.com/problems/climbing-stairs/description/
 */

/**
 * Approach: Dynamic Programming
 * I realized to climb n steps you can either take 1 or 2 steps at a time.
 * Therefore, the number of ways to reach n steps is the sum of the number of ways to reach n-1 steps and n-2 steps.
 * I used a map to store the number of ways to reach each step and then return the number of ways to reach n steps.
 * Steps:
 * 1. Initialize a map to store the number of ways to reach each step.
 * 2. If n is 1 , return 1.
 * 3. Set the number of ways to reach 0 and 1 to 1.
 * 4. For each step from 2 to n:
 *    - Set the number of ways to reach the current step to the sum of the number of ways to reach the previous two steps (n-1 and n-2).
 * 5. Return the number of ways to reach n steps.
 *
 * Time Complexity: O(n) - iterate through all elements once
 * Space Complexity: O(n) - space for the result array
 */

func climbStairs(n int) int {
	result := make(map[int]int, n+1)

	if n == 1 {
		return 1
	}

	result[0] = 1
	result[1] = 1

	for i := 2; i <= n; i++ {
		result[i] = result[i-2] + result[i-1]
	}
	return result[n]
}
