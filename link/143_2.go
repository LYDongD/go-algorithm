package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	//find the mid pointer
	mid := findMidPointer(head)

	//reverse the right half
	reverseRightHalf(mid)

	//reorder from start and mid
	reorder(head, mid)
}

func findMidPointer(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func reverseRightHalf(mid *ListNode) {
	prev, post := mid.Next, mid.Next.Next
	prev.Next = nil
	for post != nil {
		nextCursor := post.Next
		post.Next = prev
		prev = post
		post = nextCursor
	}

	mid.Next = prev
}

func reorder(head *ListNode, mid *ListNode) {
	first, preSecond, second := head, mid, mid.Next
	for first != mid {

		//delete second
		nextSecond := second.Next
		preSecond.Next = nextSecond

		//insert second
		nextFirst := first.Next
		first.Next = second
		second.Next = nextFirst
		second = nextSecond
		first = nextFirst
	}
}

func main() {
	fmt.Println("hello")
}
