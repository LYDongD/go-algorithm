package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast := head.Next
	slow := head
	var prev *ListNode

	duplicateCount := 0
	for slow != nil && fast != nil {
		if fast.Val == slow.Val {
			fast = fast.Next
			duplicateCount++
		} else {
			if duplicateCount > 0 {
				if prev != nil {
					prev.Next = fast
				} else {
					head = fast
				}
				slow = fast
				if slow != nil {
					fast = slow.Next
				}
				duplicateCount = 0
			} else {
				prev = slow
				slow = fast
				fast = fast.Next
			}
		}
	}

	if duplicateCount > 0 {

		if prev == nil {
			return prev
		}

		prev.Next = nil
	}

	return head
}

func newListNode(nums []int) *ListNode {
	var head *ListNode
	cursor := head
	for _, num := range nums {
		if cursor == nil {
			cursor = &ListNode{}
			cursor.Val = num
			head = cursor
		} else {
			cursor.Next = &ListNode{}
			cursor = cursor.Next
			cursor.Val = num
		}
	}

	return head
}

func printNode(head *ListNode) {
	cursor := head
	for cursor != nil {
		fmt.Print(cursor.Val, "->")
		cursor = cursor.Next
	}
	fmt.Println("")
}

func main() {

	node := newListNode([]int{1, 2, 3, 3, 4, 4, 5})
	node = deleteDuplicates(node)
	printNode(node)

	node2 := newListNode([]int{1, 1, 1, 1, 2})
	node2 = deleteDuplicates(node2)
	printNode(node2)
}
