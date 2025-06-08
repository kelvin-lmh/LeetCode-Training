package blind_75

/**
 * Author: Kelvin Le
 * Link:https://leetcode.com/problems/search-in-rotated-sorted-array/description/
 * Created at: 07/06/2025
 */

/**
 * Approach: Binary Search ( we check which half is sorted and check if target is in that half, if not, we search the other half)
 * Time Complexity: O(log n) - halve the search space each iteration
 * Space Complexity: O(1) - constant extra space
 */

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		}
		// left sorted portion
		if nums[l] <= nums[m] {
			if nums[m] < target || target < nums[l] {
				l = m + 1
			} else {
				r = m - 1
			}
			// right sorted portion
		} else {
			if nums[m] > target || target > nums[r] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}

	return -1

}
