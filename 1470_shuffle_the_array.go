func shuffle(nums []int, n int) []int {
	arr_size := 2 * n
	new_arr := make([]int, arr_size)
	j := 0
	for i := 0; i < n; i++ {
		new_arr[j] = nums[i]
		new_arr[j+1] = nums[i+n]
		j = j + 2
	}
	return new_arr
}