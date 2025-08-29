package blind_75

/**
 * Author: Kelvin Le
 * Problem: 242. Valid Anagram
 * Link:https://leetcode.com/problems/valid-anagram/description/
 */

/**
 * Approach: Hashmap
 * Time Complexity: O(n)
 * Space Complexity: O(1) - just have 26 characters in the map
 */

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)

	for _, ch := range s {
		count[ch]++
	}
	for _, ch := range t {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}

	return true
}
