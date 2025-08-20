package blind_75

/**
 * Author: Kelvin Le
 * Problem: 62. Unique Paths
 * Link:https://leetcode.com/problems/unique-paths/description/
 */

/**
 * Approach: Dynamic Programming
 * I use a dynamic programming approach to find the number of unique paths in a grid.
 * I maintain a 2D dp array where dp[i][j] represents the number of unique paths to reach cell (i, j).
 * The number of unique paths to reach a cell is the sum of the unique paths to reach the cell above it and the cell to the left of it.
 * I initialize the first row and first column of the dp array to 1, as there is only one way to reach any cell in the first row or first column.
 * Time Complexity: O(m * n)
 * Space Complexity: O(m * n)
 */

func uniquePaths(m int, n int) int {

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}

/**
 * Approach: Dynamic Programming
 * I only use a 1D array to optimize space complexity.
 * I maintain a dp array where dp[j] represents the number of unique paths to reach cell (i, j).
 * Instead of using a 2D array, I update the dp array in place.
 * With each cell (i,j) I update dp[j] to be the sum of dp[j] (the number of unique paths to the cell above it) and dp[j-1] (the number of unique paths to the cell to the left of it).
 * Time Complexity: O(m * n)
 * Space Complexity: O(n)
 */

// Optimize to use O(n) space
func uniquePathsOptimized(m int, n int) int {
	dp := make([]int, n)
	for j := 0; j < n; j++ {
		dp[j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[n-1]
}
