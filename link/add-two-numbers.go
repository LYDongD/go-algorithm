package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    cursor1, cursor2 := l1, l2
    var head *ListNode
    cursor := head
    carry := 0
    for cursor1 != nil ||  cursor2 != nil {

	val := 0
	if cursor1 == nil {
	    val = cursor2.Val + carry
	    cursor2 = cursor2.Next
	}else if cursor2 == nil {
	    val = cursor1.Val + carry
	    cursor1 = cursor1.Next
	}else {
	    val = cursor1.Val + cursor2.Val + carry
	    cursor1 = cursor1.Next
	    cursor2 = cursor2.Next
	}

	if val >= 10 {
	    val = val % 10
	    carry = 1
	}else {
	    carry = 0
	}

	node := &ListNode{val, nil}
	if cursor == nil {
	    cursor = node
	    head = node
	}else {
	    cursor.Next = node
	    cursor = cursor.Next
	}
    }

    if carry == 1 {
	cursor.Next = &ListNode{1, nil}
    }

    return head
}

func main() {
    printNode(addTwoNumbers(arrayToLink([]int{2,4,3}), arrayToLink([]int{5,6,4})))
    printNode(addTwoNumbers(arrayToLink([]int{1}), arrayToLink([]int{9,9})))
    printNode(addTwoNumbers(arrayToLink([]int{5,6,4}), arrayToLink([]int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1})))
}

func arrayToLink(nums []int) *ListNode {
    var head *ListNode
    cursor := head
    for _, num := range nums {
	if cursor == nil {
	    cursor = &ListNode{num, nil}
	    head = cursor
	}else {
	    cursor.Next = &ListNode{num, nil}
	    cursor = cursor.Next
	}
    }
    return head
}

func printNode(node *ListNode) {
    cursor := node
    for cursor != nil {
	fmt.Print(cursor.Val)
	if cursor.Next != nil {
	    fmt.Print("->")
	}
	cursor = cursor.Next
    }

    fmt.Println("")
}
