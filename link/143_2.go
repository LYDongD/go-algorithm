package main

import "fmt"

type ListNode struct{
    Val int
    Next *ListNode
}

func reorderList(head *ListNode)  {
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
    cursor := mid.Next
    if cursor != nil && cursor.Next != nil{
	cursor.Next.Next = cursor
	cursor = cursor.Next
    }

    mid.Next.Next = mid
}

func reorder(head *ListNode, mid *ListNode) {
    first, second := head, mid
    for first != mid {
	firstNext := first.Next
	first.Next = second
	second = second.Next
	first.Next.Next = firstNext
	first = firstNext
    }
}

func main() {
    fmt.Println("hello")
}
