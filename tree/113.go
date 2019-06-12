package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var result [][]int

func pathSum(root *TreeNode, sum int) [][]int {
    if root == nil {
	return nil
    }

    result := [][]int{}
    backtrace(root, 0, sum, []int{})
    return result
}

func backtrace(node *TreeNode, currentSum, sum int, selected []int) {
    if node.Left == nil && node.Right == nil {
        currentSum += node.Val
        tmp := make([]int, len(selected))
        copy(tmp, selected)
        tmp = append(tmp, node.Val)
        if currentSum == sum && len(tmp) > 0 {
            result = append(result, tmp)
        }

        return
    }

    tmp := make([]int, len(selected))
    copy(tmp, selected)
    tmp = append(tmp, node.Val)
    if node.Left != nil {
        backtrace(node.Left, currentSum + node.Val, sum, tmp)
    }
    if node.Right != nil {
        backtrace(node.Right, currentSum + node.Val, sum, tmp)
    }
}
func main() {
    fmt.Println("hi")
}
