package blind_75

/**
 * Author: Kelvin Le
 * Problem: 139. Word Break
 * Link:https://leetcode.com/problems/word-break/description/
 */

/**
 * Approach: Dynamic Programming
 * Time Complexity: O(n^2)
 * Space Complexity: O(n)
 */

func wordBreak(s string, wordDict []string) bool {
	n := len(s)

	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}
