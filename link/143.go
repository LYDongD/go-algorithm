package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil || head.Next.Next == nil {
	return
    }

    start := head
    for {
	prev, end := findTail(start)
	if start.Next == nil || start.Next.Next == nil {
	    return
	}

	end.Next = start.Next
	start.Next = end
	prev.Next = nil
	start = end.Next
    }
}

func findTail(start *ListNode) (*ListNode, *ListNode) {
    current := start
    var prev *ListNode
    for current.Next != nil {
	prev = current
	current = current.Next
    }
    return prev, current
}

func main() {
    fmt.Println("hi")
}
