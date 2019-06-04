package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type Queue struct {
    nodes []*LevelNode
}

type LevelNode struct {
    node *TreeNode
    level int
}

func (self *Queue)  isEmpty() bool {
    return len(self.nodes) == 0
}

func (self *Queue) offer(node *LevelNode) {
    self.nodes = append(self.nodes, node)
}

func (self *Queue) poll() *LevelNode {
    if self.isEmpty() {
	panic("queue is empty")
    }
    last := self.nodes[0]
    self.nodes = self.nodes[1:]
    return last
}


func levelOrder(root *TreeNode) [][]int {
    result := [][]int{}
    if root == nil {
	return result
    }

    queue := &Queue{nodes:[]*LevelNode{}}
    queue.offer(&LevelNode{node:root, level:1})
    currentLevel := 1
    levelList := []int{}
    for !queue.isEmpty() {
	levelNode := queue.poll()

	//add node val to level list
	if currentLevel == levelNode.level {
	    levelList = append(levelList, levelNode.node.Val)
	}else {
	    result = append(result, levelList)
	    levelList = []int{}
	    currentLevel = levelNode.level
	    levelList = append(levelList, levelNode.node.Val)
	}

	if levelNode.node.Left != nil {
	    queue.offer(&LevelNode{levelNode.node.Left, levelNode.level+1})
	}

	if levelNode.node.Right != nil {
	    queue.offer(&LevelNode{levelNode.node.Right, levelNode.level+1})
	}
    }

    //add last level
    result = append(result, levelList)
    return result
}

func main() {
    root := &TreeNode{Val:3}
    root.Left = &TreeNode{Val:9}
    root.Right = &TreeNode{Val:20}
    root.Right.Left = &TreeNode{Val:15}
    root.Right.Right = &TreeNode{Val:7}

    fmt.Println(levelOrder(root))
}
