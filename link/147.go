package main

import "fmt"
import "math"

type ListNode struct {
    Val int
    Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
	return head
    }

    //create the dummy with minest val
    dummy := &ListNode{Val:int(math.MinInt64), Next:head}
    for cursor := dummy; cursor.Next != nil; {
	//optimize
	if cursor.Val <= cursor.Next.Val {
	    cursor = cursor.Next
	    continue
	}

	//compare insert
	for orderedNode := dummy; orderedNode != cursor; orderedNode = orderedNode.Next {
	    //compare next val and that will reduce one pointer
	    if orderedNode.Next.Val > cursor.Next.Val {
		nextOrderedNode := orderedNode.Next
		orderedNode.Next = cursor.Next
		cursor.Next = cursor.Next.Next
		orderedNode.Next.Next = nextOrderedNode
		break
	    }
	}
    }

    return dummy.Next
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

func main() {
    head := makeList([]int{-1,5,3,4,0})
    node := insertionSortList(head)
    printListNode(node)
}
