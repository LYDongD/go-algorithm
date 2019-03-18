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
    Val int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {
	return head
    }

    puppy := &ListNode{0, head}
    fastCursor, slowCursor := puppy, puppy
    i := 0
    for fastCursor != nil  {
	fastCursor = fastCursor.Next
	i++
	if i == n {
	    break
	}
    }

    if fastCursor == nil {
	return nil
    }

    for fastCursor.Next != nil {
	slowCursor = slowCursor.Next
	fastCursor = fastCursor.Next
    }

    if slowCursor.Next == head {
	head = head.Next
    }else {
	slowCursor.Next = slowCursor.Next.Next
    }

    puppy.Next = nil

    return head
}
