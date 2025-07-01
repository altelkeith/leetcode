func transformArray(nums []int) []int {
	arr_size := len(nums)
	for i := 0; i < arr_size; i++ {
		if nums[i]%2 == 0 {
			nums[i] = 0
		} else {
			nums[i] = 1
		}
	}
	slices.Sort(nums)
	return nums
}