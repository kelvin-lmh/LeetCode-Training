package blind_75

func productExceptSelf(nums []int) []int {
	results := make([]int, len(nums))
	for i := range results {
		results[i] = 1
	}
	prefixProduct := 1
	for i := 0; i < len(nums); i++ {
		results[i] = prefixProduct
		prefixProduct *= nums[i]
	}
	suffixProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		results[i] *= suffixProduct
		suffixProduct *= nums[i]
	}
	return results
}
