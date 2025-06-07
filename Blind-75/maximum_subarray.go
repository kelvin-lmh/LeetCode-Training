package blind_75

func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0
	for i, _ := range nums {
		sum += nums[i]
		if nums[i] > sum {
			sum = nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
