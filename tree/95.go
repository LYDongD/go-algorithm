package main

import "fmt"


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
    if n < 1 {
	return []*TreeNode{}
    }

    return generateByDC(1,n)
}

func generateByDC(start, end int) []*TreeNode {
    result := []*TreeNode{}
    if start > end {
	result = append(result, nil)
	return result
    }

    if start == end {
	result = append(result, &TreeNode{Val:start})
	return result
    }

    for i := start; i <= end; i++ {
	leftList := generateByDC(start, i - 1)
	rightList := generateByDC(i+1, end)
	for k := 0; k < len(leftList); k++ {
	    for t := 0; t < len(rightList); t++ {
		root := &TreeNode{Val:i}
		root.Left = leftList[k]
		root.Right = rightList[t]
		result = append(result, root)
	    }
	}

    }

    return result
}


func main() {
    fmt.Println(generateTrees(3))
}
