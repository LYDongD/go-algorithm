package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	cache := make(map[*ListNode]bool)
	cursor := headA
	for cursor != nil {
		cache[cursor] = true
		cursor = cursor.Next
	}

	cursor = headB
	for cursor != nil {
		if cache[cursor] != false {
			return cursor
		}

		cursor = cursor.Next
	}

	return nil
}
