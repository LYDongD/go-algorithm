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

    //fast-slow pointer to find the meet node
    slow, fast := head, head

    for {
	if fast == nil || fast.Next == nil {
	    return nil
	}
	slow = slow.Next
	fast = fast.Next.Next
	if fast == slow {
	    break
	}
    }

    //move from start and meet node, find next meet
    first, second := head, fast
    for first != second {
	first = first.Next
	second = second.Next
    }

    return first
}

func main() {
	fmt.Println("haha")
}
