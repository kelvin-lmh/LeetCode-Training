package blind_75

import "sort"

/**
 * Author: Kelvin Le
 * Problem: 15. 3Sum
 * Link:https://leetcode.com/problems/3sum/description/
 * Created at: 08/06/2025
 */

/**
 * Approach: Two Pointers
 * Time Complexity: O(n^2) - nested loops
 * Space Complexity: O(1) - constant extra space
 * Steps:
 * 1. Sort the array to handle duplicates and use two pointers effectively.
 * 2. Iterate through the array, using the current element as a fixed point.
 * 3. For each fixed point, use two pointers (left and right) to find pairs that sum to the negative of the fixed point.
 * 4. Skip duplicates to avoid counting the same triplet multiple times.
 * 5. Return the list of unique triplets that sum to zero.
 */

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	for i, n := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			threeSum := n + nums[l] + nums[r]
			if threeSum > 0 {
				r--
			} else if threeSum < 0 {
				l++
			} else {
				result = append(result, []int{n, nums[l], nums[r]})
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}

	}
	return result
}
