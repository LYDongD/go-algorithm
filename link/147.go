package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
	return head
    }

    dummy := &ListNode{Next:head}
    prevCursor, cursor := head, head.Next
    for (cursor != nil) {
	nextCursor := cursor.Next
	prevNode, node := dummy, dummy.Next
	for node != cursor {
	    if node.Val > cursor.Val {
		prevCursor.Next = nextCursor
		cursor.Next = node
		prevNode.Next = cursor
		break
	    }

	    prevNode = node
	    node = node.Next
	}

	cursor = nextCursor
    }

    return dummy.Next
}

func main() {
    fmt.Println("hah")
}
