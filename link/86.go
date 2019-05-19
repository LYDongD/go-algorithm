package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	//init pointer
	dummy := &ListNode{0, head}
	cursor, pivot := head, dummy
	var prev *ListNode

	//find and move node less than target x
	for cursor != nil {
		if cursor.Val < x {
			current := cursor
			cursor = current.Next

			if pivot.Next == current { //no need to delete, just update prev
				prev = current
			} else { //delete and insert after pivot
				prev.Next = current.Next
				current.Next = pivot.Next
				if pivot.Next == head { //insert into head，reset head
					head = current
				}
				pivot.Next = current
			}

			//update pivot
			pivot = pivot.Next
		} else {
			prev = cursor
			cursor = cursor.Next
		}
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
	node := newListNode([]int{1, 4, 3, 2, 5, 2})
	node = partition(node, 5)
	printNode(node)
}
