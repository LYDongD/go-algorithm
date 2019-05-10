package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{0, head}
	last := dummy
	current := head

	for current != nil && current.Next != nil {
		next := current.Next
		current.Next = next.Next
		next.Next = current
		last.Next = next

		last = current
		current = current.Next
	}

	return dummy.Next
}

func main() {
	fmt.Println("")
}
