func numIdenticalPairs(nums []int) int {
	arr_size := len(nums)
	counter := 0
	for i := 0; i < arr_size; i++ {
		for j := 0; j < arr_size; j++ {
			if nums[i] == nums[j] && i < j {
				counter = counter + 1
			}
		}
	}
	return counter
}