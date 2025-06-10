package blind_75

/**
* Author: Kelvin Le
 * Problem: 191. Number of 1 Bits
* Link:https://leetcode.com/problems/number-of-1-bits/description/
* Created at: 10/06/2025
*/

/**
 * Approach: Bit Manipulation
 * - Use bitwise operations to count the number of 1 bits in the binary representation of an integer.
 * - The idea is to repeatedly check the least significant bit and right shift the integer until it becomes zero.
 * Time Complexity: O(log n) - iterating through the bits of the integer once (because the number of bits in an integer is log(n) where n is the value of the integer).
 * Space Complexity: O(1) - constant extra space
 */

func hammingWeight(n int) int {
	count := 0
	for n != 0 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}
	return count
}
