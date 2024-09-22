package main

import "fmt"

func binarySearch(orderedNums []int, goal int) int {
	currentMinIndex := 0
	currentMaxIndex := len(orderedNums) - 1
	// check if goal is out of range
	if goal < orderedNums[currentMinIndex] || goal > orderedNums[currentMaxIndex] {
		return -1
	}
	// divide the current range in half until index is found or we have searched
	// the entire range
	for currentMaxIndex <= currentMinIndex {
		indexOfInterest := currentMinIndex + (currentMaxIndex-currentMinIndex)/2
		switch {
		case goal < orderedNums[indexOfInterest]:
			currentMaxIndex = indexOfInterest - 1
		case goal > orderedNums[indexOfInterest]:
			currentMinIndex = indexOfInterest + 1
		default:
			return indexOfInterest
		}
	}
	return -1
}

func main() {
	orderedNums := []int{1, 3, 5, 7, 9}
	goal := 3
	index := binarySearch(orderedNums, goal)
	fmt.Printf("The index of %d is %d \n", goal, index)
}
