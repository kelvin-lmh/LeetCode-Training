package blind_75

/**
 * Author: Kelvin Le
 * Problem: 3. Longest Substring Without Repeating Characters
 * Link:https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
 */

/**
 * Approach: Sliding Window
 * Steps:
 * 1. Initialize a map to store the last seen index of each character.
 * 2. Initialize a variable `maxLen` to keep track of the maximum length of the longest substring without repeating characters.
 * 3. Initialize a variable `left` to keep track of the left index of the current substring.
 * 4. Iterate through the string `s` from index 0 to the length of the string.
 * 5. If the current character is already in the map, update `left` to the index of the last seen occurrence of the character plus 1.
 * 6. Add the current character and its index to the map.
 * 7. Update `maxLen` if the current substring length is greater than the previous maximum length.
 * 8. Return the maximum length.
 *
 * Time Complexity: O(n) - iterate through all elements once
 * Space Complexity:  O(n)
 */
func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	left := 0
	lastSeen := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		if idx, found := lastSeen[s[right]]; found && idx >= left {
			left = idx + 1
		}
		lastSeen[s[right]] = right
		if maxLen < right-left+1 {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
