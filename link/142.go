package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	start := head
	for {
		slow, fast := start, start.Next
		if fast == slow {
			return start
		}

		for fast != slow {

			if fast == nil || fast.Next == nil || fast.Next.Next == nil {
				return nil
			}

			if fast == start || fast.Next == start {
				return start
			}

			fast = fast.Next.Next
			slow = slow.Next
		}
		start = start.Next
	}

	return nil
}

func main() {
	fmt.Println("haha")
}
