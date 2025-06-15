package blind_75

import (
	"sort"
)

/**
 * Author: Kelvin Le
 * Problem: 300. Longest Increasing Subsequence
 * Link:https://leetcode.com/problems/longest-increasing-subsequence/description/
 */

/**
 * Approach: Dynamic Programming
 * I used a dp array to store the length of the longest increasing subsequence ending at each index in nums.
 * I then iterate through the array and update the dp array with the maximum length of the longest increasing subsequence ending at each index.
 * Time Complexity: O(n^2)
 * Space Complexity: O(n)
 */

func lengthOfLIS(nums []int) int {
	maxLength := 1
	dp := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
				maxLength = max(maxLength, dp[i])
			}
		}
	}

	return maxLength
}

/**
 * Approach: Binary Search
 * I used a sub array to store the longest increasing subsequence ending at each index in nums.
 * I then iterate through the array and update the sub array with the longest increasing subsequence ending at each index.
 * Time Complexity: O(n log n)
 * Space Complexity: O(n)
 */

func lengthOfLISOptimized(nums []int) int {
	sub := []int{}

	for _, x := range nums {
		i := sort.SearchInts(sub, x)
		if i == len(sub) {
			sub = append(sub, x)
		} else {
			sub[i] = x
		}
	}

	return len(sub)
}
