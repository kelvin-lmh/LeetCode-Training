package blind_75

/**
 * Author: Kelvin Le
 * Link:https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
 * Created at: 07/06/2025
 */

/**
 * Approach: Linear Scan
 * Time Complexity: O(n) - iterate through all elements once
 * Space Complexity: O(1) - constant extra space
 */
func findMin(nums []int) int {
	m := nums[0]
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m
}

/**
 * Approach: Binary Search – Find the half that contains the minimum
 * This approach is applicable when the input array is a rotated sorted array with no duplicates.
 *
 * Time Complexity: O(log n) – Each iteration halves the search space.
 * Space Complexity: O(1) – Constant extra space.
 */

func findMinBinarySearch(nums []int) int {
	res := nums[0]
	l := 0
	r := len(nums) - 1

	for l <= r {
		// If the current subarray is sorted, take the leftmost as minimum
		if nums[l] <= nums[r] {
			res = min(res, nums[l])
			break
		}

		m := (l + r) / 2
		res = min(res, nums[m])

		// Left half is sorted, so minimum must be in right half
		if nums[m] >= nums[l] {
			l = m + 1
		} else { // Right half is unsorted, so minimum is there
			r = m - 1
		}
	}

	return res
}
