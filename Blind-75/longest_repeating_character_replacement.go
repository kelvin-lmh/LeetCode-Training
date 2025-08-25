package blind_75

/**
 * Author: Kelvin Le
 * Problem: 424. Longest Repeating Character Replacement
 * Link:https://leetcode.com/problems/longest-repeating-character-replacement/description/
 */

/**
 * Approach: Sliding Window
 * Steps:
 * 1. Initialize a map to store the count of each character in the current window.
 * 2. Initialize a variable `left` to keep track of the left index of the current window.
 * 3. Initialize a variable `maxLen` to keep track of the maximum length of the longest substring without repeating characters.
 * 4. Initialize a variable `maxCount` to keep track of the maximum count of any character in the current window.
 * 5. Iterate through the string `s` from index 0 to the length of the string.
 * 6. Increment the count of the current character in the map.
 * 7. If the current character's count is greater than the current `maxCount`, update `maxCount`.
 * 8. If the current window length minus the maximum count of any character is greater than `k`, decrement the count of the character at the left index and increment `left`.
 * 9. If the current window length greater than the previous `maxLen`, update `maxLen`.
 * 10. Return the `maxLen`.
 * Time Complexity: O(n)
 * Space Complexity: O(1)
 */

func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	left, maxLen, maxCount := 0, 0, 0

	for right := 0; right < len(s); right++ {
		count[s[right]]++
		if count[s[right]] > maxCount {
			maxCount = count[s[right]]
		}
		if (right-left+1)-maxCount > k {
			count[s[left]]--
			left += 1
		}
		if (right-left+1)-maxLen > 0 {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
