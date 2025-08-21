package blind_75

/**
 * Author: Kelvin Le
 * Problem: 55. Jump Game
 * Link:https://leetcode.com/problems/jump-game/description/
 */

/**
 * Approach: Greedy
 * I use a greedy approach to determine if I can reach the last index of the array.
 * Steps:
 * 1. Initialize a variable `farthest` to keep track of farthest index reachable from the current index.
 * 2. Iterate through the array .
 * 3. If at any point the current index exceeds `farthest`, return false.
 * 4. Update `farthest` to be the maximum of its current value and the sum of the current index and the jump length at that index.
 * 5. If `farthest` reaches or exceeds the last index, return true.
 * Time Complexity: O(n) - single loop
 * Space Complexity: O(1) - constant extra space
 */

func canJump(nums []int) bool {
	n := len(nums)
	farthest := 0
	for i := 0; i <= n-1; i++ {
		if i > farthest {
			return false
		}
		farthest = max(farthest, i+nums[i])
		if farthest >= n-1 {
			return true
		}
	}
	return true
}
