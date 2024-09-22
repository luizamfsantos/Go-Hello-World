package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Merge the two lists into one sorted list.
// The list should be made by splicing together the nodes of the first two lists.

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	for i := 0; i < len(list1); i++ {
// 		for j := 0; j < len(list2); j++ {

// 		}
// 	}
// }

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{}}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{}}}}
	fmt.Println(list1, list2)
}
