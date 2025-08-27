package blind_75

/**
 * Author: Kelvin Le
 * Problem: 76. Minimum Window Substring
 * Link:https://leetcode.com/problems/minimum-window-substring/description/
 */

/**
 * Approach: Sliding Window
 * Count the frequency of characters in t, use a map[byte]int to store the number of times each character is needed.
 * Use two pointers (left, right) to browse on s:
 * - right expands the window → add more characters.
 * - left collapses the window → when there are enough characters in t.
 * If the current window contains enough characters and is shorter than the previous result → update the result.
 * Steps:
 * 1. Initialize a map to store the frequency of characters in t.
 * 2. Initialize a map to store the frequency of characters in the current window.
 * 3. Iterate through the string s using two pointers, left and right.
 * 4. Initialize a variable have to keep track of the number of characters in t and their frequency in the current window.
 * 5. Increment the frequency of the character at the right pointer in the window map.
 * 6. If the character at the right pointer is in the t map and its frequency in the window map is equal to the frequency in t, increment the have variable.
 * 7. If the have variable is equal to the length of t, check if the current window is smaller than the result window. If it is, update the result window.
 * 8. Decrement the frequency of the character at the left pointer in the window map.
 * 8. If the character at the left pointer is in the t map and its frequency in the window map is less than the frequency in t, decrement the have variable.
 * 10. If the have variable is equal to the length of t, check if the current window is smaller than the need map. If it is, update the need map.
 * 11. Return the result window.
 * Time Complexity: O(m + n)
 * Space Complexity: O(1) - because there are only 26 English characters
 */

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	// count the frequency of elements in t
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	// sliding window
	window := make(map[byte]int)
	left := 0
	have, needCount := 0, len(need)
	res, resLen := [2]int{-1, -1}, len(s)+1

	for right := 0; right < len(s); right++ {
		window[s[right]]++

		if need[s[right]] > 0 && need[s[right]] == window[s[right]] {
			have++
		}

		for have == needCount {
			if right-left+1 < resLen {
				resLen = right - left + 1
				res = [2]int{left, right}
			}

			window[s[left]]--
			if need[s[left]] > 0 && window[s[left]] < need[s[left]] {
				have--
			}
			left++
		}
	}
	if resLen == len(s)+1 {
		return ""
	}
	return s[res[0] : res[1]+1]
}
