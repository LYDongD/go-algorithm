package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type Queue struct {
    nodes []*TreeNode
    levels []int
}

func (this *Queue) Offer(node *TreeNode, level int) {
    this.nodes = append(this.nodes, node)
    this.levels = append(this.levels, level)
}

func (this *Queue) IsEmpty() bool {
    return len(this.nodes) == 0
}

func (this *Queue) Take() (*TreeNode, int)  {
    if this.IsEmpty() {
	panic("queue is empty")
    }

    head := this.nodes[0]
    level := this.levels[0]
    this.nodes = this.nodes[1:]
    this.levels = this.levels[1:]
    return head, level
}

func levelOrder(root *TreeNode) [][]int {
    result := [][]int{}
    if root == nil {
	return result
    }

    queue := &Queue{[]*TreeNode{}, []int{}}
    queue.Offer(root, 0)
    currentLevel := 0
    currentLevelResult := []int{}
    left2Right := true
    for !queue.IsEmpty() {
	node, level := queue.Take()
	if level == currentLevel {
	    if left2Right {
		currentLevelResult = append(currentLevelResult, node.Val)
	    }else {
		currentLevelResult = append([]int{node.Val}, currentLevelResult...)
	    }
	}else {
	    left2Right = !left2Right
	    result = append(result, currentLevelResult)
	    currentLevelResult = []int{node.Val}
	    currentLevel = level
	}

	if node.Left != nil {
	    queue.Offer(node.Left, level+1)
	}

	if node.Right != nil {
	    queue.Offer(node.Right, level+1)
	}
    }

    if len(currentLevelResult) > 0 {
	result = append(result, currentLevelResult)
    }

    return result
}

func main() {
    fmt.Println("GO")
}
