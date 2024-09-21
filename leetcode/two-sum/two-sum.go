package main

import "fmt"

// Given an array of integers nums and an integer target,
// return the indices of the two numbers such that
// they add up to target.
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		leftOver := target - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == leftOver {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	answer1 := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Printf("Answer: %d\n", answer1) // [0, 1]

	answer2 := twoSum([]int{3, 2, 4}, 6)
	fmt.Printf("Answer: %d\n", answer2) // [1, 2]

	answer3 := twoSum([]int{3, 3}, 6)
	fmt.Printf("Answer: %d\n", answer3) // [0, 1]
}
