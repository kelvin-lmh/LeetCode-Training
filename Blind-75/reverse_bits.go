package blind_75

/**
 * Author: Kelvin Le
 * Problem: 190. Reverse Bits
 * Link:https://leetcode.com/problems/reverse-bits/description/
 */

/**
 * Approach: Bit Manipulation
 * Time Complexity: O(1)
 * Space Complexity: O(1)
 * Steps:
 * 1. Initialize a result variable to zero.
 * 2. For each bit in the input number:
 *    - Shift the result to the left by one position.
 *    - Add the rightmost bit of the input number to the result.
 *    - Right shift the input number by one position.
 * 3. Return the result.
 */

func reverseBits(num uint32) uint32 {
	var result uint32 = 0
	for i := 0; i < 32; i++ {
		bit := num & 1
		result = (result << 1) | bit
		num >>= 1
	}
	return result
}
