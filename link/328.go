package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cursor, handle := head, head

	for cursor != nil {
		prev := cursor.Next
		if prev == nil {
			return head
		}
		cursor = prev.Next
		if cursor == nil {
			return head
		}

		//delete odd cursor
		prev.Next = cursor.Next

		//insert
		cursor.Next = handle.Next
		handle.Next = cursor
		handle = cursor

		//reset
		cursor = prev

		printerLink(head)
	}

	return head
}

func generateLink(nums []int) *ListNode {
	var head *ListNode

	var node *ListNode
	for _, num := range nums {
		if node == nil {
			node = &ListNode{}
			node.Val = num
			head = node
		} else {
			nextNode := &ListNode{}
			nextNode.Val = num
			node.Next = nextNode
			node = nextNode
		}
	}

	return head
}

func printerLink(head *ListNode) {
	node := head
	for node != nil {
		fmt.Print(node.Val, "->")
		node = node.Next
	}
	fmt.Print("NULL")
	fmt.Println()
}

func main() {
	node := generateLink([]int{1, 2, 3, 4, 5})
	printerLink(node)
	printerLink(oddEvenList(node))
	node2 := generateLink([]int{2, 1, 3, 5, 6, 4, 7})
	printerLink(node2)
	printerLink(oddEvenList(node2))
}
