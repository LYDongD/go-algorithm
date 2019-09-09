package main

import "fmt"

var result int

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {

    if root == nil || k < 1 {
	panic("parameters illegal")
    }

    count := 0
    findKInorder(root, k, &count)

    return result
}

func findKInorder(node *TreeNode, k int, count *int){
    if node == nil {
	return
    }

    findKInorder(node.Left, k, count)
    *count = *count+1
    if *count == k {
	result = node.Val
	return
    }
    findKInorder(node.Right, k, count)
}


func main() {
    fmt.Println("hello")
}
