package main

import "fmt"

func binarySearch(orderedNums []int, goal int) int {
	currentMinIndex := 0
	currentMaxIndex := len(orderedNums) - 1
	// check if goal is min or max
	switch goal {
	case orderedNums[currentMinIndex]:
		return currentMinIndex
	case orderedNums[currentMaxIndex]:
		return currentMaxIndex
	}
	// divide the current range in half until index is found or we have searched
	// the entire range
	for {
		indexOfInterest := int((currentMaxIndex - currentMinIndex) / 2)
		switch {
		case goal < orderedNums[indexOfInterest]:
			currentMaxIndex = indexOfInterest
		case goal > orderedNums[indexOfInterest]:
			currentMinIndex = indexOfInterest
		default:
			return indexOfInterest
		}
		if currentMaxIndex == currentMinIndex {
			break
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
