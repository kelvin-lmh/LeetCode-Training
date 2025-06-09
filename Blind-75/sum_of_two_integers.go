package blind_75

/**
 * Author: Kelvin Le
 * Problem: 371. Sum of Two Integers
 * Link:https://leetcode.com/problems/sum-of-two-integers/description/
 */

/**
 * Approach: Bit Manipulation
 * - Use bitwise operations to calculate the sum without using the '+' operator.
 * - The idea is to use XOR to add the numbers without carrying, and & followed by a left shift to calculate the carry.
 * Time Complexity: O(1) - constant time
 * Space Complexity: O(1) - constant extra space
 * Steps:
 * 1. Use XOR to add the two numbers without carrying.
 * 2. Use & to find the carry bits, and shift them left by one position.
 * 3. Repeat steps 1 and 2 until there is no carry left.
 * 4. Return the final sum.
 */

func getSum(a int, b int) int {
	for b != 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
	}
	return a
}
