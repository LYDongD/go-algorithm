package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

var left *ListNode
var stop bool

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m >= n || head == nil {
		return head
	}

	//two pointers
	left = head
	stop = false

	//move left & right to the start and the end of sublist
	moveAndReverse(head, m, n)
	return head
}

func moveAndReverse(right *ListNode, m, n int) {
	if n == 1 {
		return
	}

	right = right.Next
	if m > 1 {
		left = left.Next
	}

	moveAndReverse(right, m-1, n-1)

	if right == left || right.Next == left {
		//todo
		stop = true
	}

	//swap
	if !stop {
		tmp := right.Val
		right.Val = left.Val
		left.Val = tmp
		//left forward & right backward while stack frame pop
		left = left.Next
	}
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
	head := newListNode([]int{1, 2, 3, 4, 5})
	head = reverseBetween(head, 2, 4)
	printNode(head)
}
