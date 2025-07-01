package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	past_max := 0
	current_max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == '1' {
			current_max = current_max + 1
		}
		if nums[i] == '0' {
			past_max = current_max
			current_max = 0
		}
	}
	if current_max > past_max {
		fmt.Println(current_max)
		return current_max
	} else {
		fmt.Println(past_max)
		return past_max
	}

}

func main() {
	numms := []int{1, 1, 0, 1, 1, 1}
	findMaxConsecutiveOnes(numms)
}
