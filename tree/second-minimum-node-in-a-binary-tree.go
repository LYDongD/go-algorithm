package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
    min := root.Val
    queue := &ListQueue{}
    queue.offer(root)
    secondMin := -1
    for !queue.isEmpty() {
	node := queue.poll().(*TreeNode)

	if (secondMin < 0 || node.Val < secondMin) && node.Val > min {
	    secondMin = node.Val
	}else if node.Val == min {
	    if node.Left != nil {
		queue.offer(node.Left)
	    }

	    if node.Right != nil {
		queue.offer(node.Right)
	    }
	}
    }

    if secondMin > min {
	return secondMin
    }

    return -1
}

type Element interface{}

type Queue interface {
    offer(e Element)
    poll() Element
    isEmpty() bool
}

type ListQueue struct {
    elements []Element
}

func (self *ListQueue) isEmpty() bool {
    return len(self.elements) == 0
}

func (self *ListQueue) offer(e Element) {
    self.elements = append(self.elements, e)
}

func (self *ListQueue) poll() Element {
    if self.isEmpty() {
	panic("queue is empty")
    }

    head := self.elements[0]
    self.elements = self.elements[1:]
    return head
}

func main() {
    node := &TreeNode{Val:5}
    node.Left = &TreeNode{Val:8}
    node.Right = &TreeNode{Val:5}
    fmt.Println(findSecondMinimumValue(node))
}
