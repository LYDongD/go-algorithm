package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
	return head
    }

    dummy := &ListNode{0, head}
    fast, slow := dummy, dummy
    for fast.Next != nil && fast.Next.Next != nil {
	fast = fast.Next.Next
	slow = slow.Next
    }

    return slow.Next
}
