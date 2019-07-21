package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
	return head
    }

    cursor := head
    for cursor.Next != nil{
	cursor = cursor.Next
    }
    start, _ := quickSort(head, cursor)
    return start
}

func quickSort(start *ListNode, end *ListNode) (*ListNode, *ListNode){

    if start == nil || end == nil || start == end {
	return start, end
    }

    pivot, start, end := partion(start, end)
    leftStart, leftEnd := quickSort(start, pivot)
    rightStart, rightend := quickSort(pivot.Next, end)
    leftEnd.Next = rightStart
    return leftStart, rightend
}

func partion(start *ListNode, end *ListNode) (*ListNode, *ListNode, *ListNode) {
    pivot := start
    head, tail := pivot, pivot
    for cursor := start.Next; cursor != nil; {
	nextCursor := cursor.Next
	if cursor.Val < pivot.Val {
	       tail.Next = nextCursor
	       cursor.Next = head
	       head = cursor
	}else {
	    tail = cursor
	}
	if cursor == end {
	    break
	}
	cursor = nextCursor
    }

    return pivot, head, tail
}

func makeList(nums []int) *ListNode {
    head := &ListNode{Val : nums[0]}
    cursor := head
    for i := 1; i< len(nums); i++  {
        cursor.Next = &ListNode{Val: nums[i]}
        cursor = cursor.Next
    }
    return head
}

func printListNode(head *ListNode) {
    cursor := head
    for cursor != nil {
        fmt.Print(cursor.Val)
        if cursor.Next != nil {
            fmt.Print("->")
        }
        cursor = cursor.Next
    }
    fmt.Println("")
}

func findTail(head *ListNode) *ListNode {
    cursor := head
    for cursor.Next != nil{
        cursor = cursor.Next
    }
    return cursor
}

func main() {
    //node := makeList([]int{4,2,1,3})
    node := makeList([]int{-1,5,3,4,0})
    //_, start, _ := partion(node, findTail(node))
    printListNode(sortList(node))
}
