package blind_75

/**
 * Author: Kelvin Le
 * Problem: 268. Missing Number
 * Link:https://leetcode.com/problems/missing-number/description/
 */

/**
* Approach: Sorting
* Time Complexity: O(n log(n))
* Space Complexity: O(1)
* Steps:
* 1. Sort the array in ascending order.
* 2. If the first element is not 0, return 0.
* 3. For each element in the sorted array, check if the difference between the current element and the previous element is greater than 1.
	 If it is, return the current element minus 1.
* 4. Return the last element of the sorted array plus 1.
*/

import "sort"

func missingNumber(nums []int) int {
	sort.Ints(nums)
	if nums[0] != 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {

		if nums[i]-nums[i-1] > 1 {
			return nums[i] - 1
		}
	}
	return nums[len(nums)-1] + 1
}

/**
* Approach: Math (Gauss' Formula)
* Time Complexity: O(n)
* Space Complexity: O(1)
* Steps:
* 1. Calculate the expected sum of all numbers from 0 to n, where n is the length of the array (n * (n + 1) / 2).
* 2. Calculate the sum of all elements in the array.
* 3. Return the difference between the two sums.
 */

func missingNumberOptimized(nums []int) int {
	expectedSum := (len(nums) * (len(nums) + 1)) / 2
	curentSum := 0

	for i := 0; i < len(nums); i++ {
		curentSum += nums[i]
	}
	return expectedSum - curentSum
}
