package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {

	if head == nil {
		return nil
	}

	// transtate link to array
	array := translateToArray(head)

	// build tree recursively
	return buildBST(array, 0, len(array)-1)
}

func translateToArray(head *ListNode) []int {
	result := []int{}
	cursor := head
	for cursor != nil {
		result = append(result, cursor.Val)
		cursor = cursor.Next
	}

	return result
}

func buildBST(array []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := start + (end-start)/2
	node := &TreeNode{Val: array[mid]}
	node.Left = buildBST(array, start, mid-1)
	node.Right = buildBST(array, mid+1, end)

	return node
}

func constructLink(array []int) *ListNode {

	root := &ListNode{Val: array[0]}
	cursor := root
	for i := 1; i < len(array); i++ {
		cursor.Next = &ListNode{Val: array[i]}
		cursor = cursor.Next
	}

	return root
}

func main() {

	//construct a link
	head := constructLink([]int{-10, -3, 0, 5, 9})

	root := sortedListToBST(head)
	fmt.Println(root, root.Left, root.Right)
}
