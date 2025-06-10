package blind_75

/**
* Author: Kelvin Le
* Problem: 11. Container With Most Water
* Link:https://leetcode.com/problems/container-with-most-water/description/
* Created at: 07/06/2025
 */

/**
 * Approach: Brute Force
 * Time Complexity: O(n2) - iterate through all elements twice
 * Space Complexity: O(1) - constant extra space
 */

func maxArea(height []int) int {
	max := 0
	for i := 0; i <= len(height)-1; i++ {
		for j := i + 1; j <= len(height)-1; j++ {
			area := min(height[i], height[j]) * (j - i)
			if area > max {
				max = area
			}
		}
	}
	return max
}

/**
 * Approach: Two Pointers
 * - Start with two pointers, one at the beginning and one at the end of the array.
 * Time Complexity: O(n) - iterate through all elements once
 * Space Complexity: O(1) - constant extra space
 * Steps:
 * 1. Initialize two pointers, left at the start and right at the end of the array.
 * 2. Calculate the area formed by the lines at these two pointers.
 * 3. Move the pointer pointing to the shorter line inward, as this will potentially increase the area.
 * 4. Repeat steps 2-3 until the two pointers meet.
 * 5. Return the maximum area found.
 */

func maxAreaTwoPointers(height []int) int {
	max := 0
	l := 0
	r := len(height) - 1

	for l < r {
		amount := min(height[l], height[r]) * (r - l)
		if amount > max {
			max = amount
		}
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return max
}
