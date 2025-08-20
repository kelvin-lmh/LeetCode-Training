package blind_75

/**
 * Author: Kelvin Le
 * Problem: 91. Decode Ways
 * Link:https://leetcode.com/problems/decode-ways/description/
 */

/**
 * Approach: Dynamic Programming
 * I use a dynamic programming approach to count the number of ways to decode the message.
 * I maintain a dp array where dp[i] represents the number of ways to decode the substring s[0:i].
 * It have two cases:
 * 1. If the current character is not '0', it can be decoded as a single character.
 * 2. If the last two characters form a valid letter (between '10' and '26'), they can be decoded as a pair.
 * I iterate through the string and update the dp array accordingly.
 * Time Complexity: O(n) - iterate through all elements once
 * Space Complexity: O(n) - dp array to store results for problems
 */

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 1
	if s[0] != '0' {
		dp[1] = 1
	} else {
		dp[1] = 0
	}

	for i := 2; i <= n; i++ {
		// case one number
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		// case two number
		two := ((s[i-2]-'0')*10 + (s[i-1] - '0'))
		if 10 <= two && two <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}
