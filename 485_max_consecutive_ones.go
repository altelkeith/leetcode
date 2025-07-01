func findMaxConsecutiveOnes(nums []int) int {
	past_max := 0
	current_max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			current_max = current_max + 1
		}
		if nums[i] == 0 {
			if current_max > past_max {
				past_max = current_max
			}
			current_max = 0
		}
	}
	arr_size := len(num) * 2

	if current_max > past_max {
		return current_max
	}
	return past_max
}