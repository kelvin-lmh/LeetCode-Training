package blind_75

/**
 * Author: Kelvin Le
 * Problem: 338. Counting Bits
 * Link:https://leetcode.com/problems/counting-bits/description/
 */

/**
 * Approach: Iterate through each number from 0 to n and count the number of 1 bits in its binary representation.
 * - For each number, use bitwise operations to count the number of 1 bits.
 * Time Complexity: O(n * log n) - iterating through each number and counting bits (log n is the number of bits in the binary representation of n)
 * Space Complexity: O(n) - space for the result array
 * Steps:
 * 1. Initialize an empty result array of size n+1.
 * 2. For each number from 0 to n:
 *    - Initialize a count variable to zero.
 *    - Use a temporary variable to check each bit of the number:
 *      - While the temporary variable is not zero:
 *        - If the least significant bit is 1 (using bitwise AND with 1), increment the count.
 *        - Right shift the temporary variable by one position to check the next bit.
 *    - Append the count to the result array.
 * 3. Return the result array.
 */

func countBits(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		count := 0
		temp := i
		for temp != 0 {
			if temp&1 == 1 {
				count++
			}
			temp >>= 1
		}
		result[i] = count
	}
	return result
}

/**
 * Approach: Optimized Bit Manipulation
 * - Use a dynamic programming approach where the number of 1 bits in a number i can be derived from the number of 1 bits in i >> 1 (i divided by 2) and the least significant bit (i&1).
 * - The number of 1 bits in i is equal to the number of 1 bits in i >> 1 plus 1 if the least significant (Rightmost) bit is 1.
 * Time Complexity: O(n) - iterating through each number once
 * Space Complexity: O(n) - space for the result array
 */

func countBitsOptimized(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		result[i] = result[i>>1] + (i & 1)
	}
	return result
}
