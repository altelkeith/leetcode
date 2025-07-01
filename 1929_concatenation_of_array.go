func getConcatenation(nums []int) []int {
	old_size := len(nums)
	arr_size := old_size * 2
	ans := make([]int, arr_size)
	for i := 0; i < old_size; i++ {
		ans[i] = nums[i]
		ans[i+old_size] = nums[i]
	}
	return ans
}