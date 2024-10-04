package main

import "fmt"

// the three dots before int means an arbitrary number
// of ints as arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	// for _, num := range nums {
	// 	total += num
	// }
	for j := 0; j < len(nums); j++ {
		total += nums[j]
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
