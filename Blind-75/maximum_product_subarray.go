package blind_75

func maxProduct(nums []int) int {
	result := nums[0]
	curMax := nums[0]
	curMin := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			curMax, curMin = curMin, curMax
		}

		curMax = max(curMax*nums[i], nums[i])
		curMin = min(curMin*nums[i], nums[i])
		result = max(result, curMax)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
