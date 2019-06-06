package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type LevelNode struct {
	node  *TreeNode
	level int
}

type Queue struct {
	nodes []*LevelNode
}

func (self *Queue) isEmpty() bool {
	return len(self.nodes) == 0
}

func (self *Queue) offer(node *LevelNode) {
	self.nodes = append(self.nodes, node)
}

func (self *Queue) poll() *LevelNode {
	if self.isEmpty() {
		panic("queue is empty")
	}

	head := self.nodes[0]
	self.nodes = self.nodes[1:]
	return head
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := &Queue{nodes: []*LevelNode{}}
	queue.offer(&LevelNode{root, 1})
	currentLevel := 1
	levelList := []int{}
	for !queue.isEmpty() {
		levelNode := queue.poll()
		if levelNode.level == currentLevel {
			if currentLevel%2 == 1 {
				levelList = append(levelList, levelNode.node.Val)
			} else {
				levelList = append([]int{levelNode.node.Val}, levelList...)
			}
		} else {
			result = append(result, levelList)
			currentLevel = levelNode.level
			levelList = []int{}
			levelList = append(levelList, levelNode.node.Val)
		}

		if levelNode.node.Left != nil {
			queue.offer(&LevelNode{levelNode.node.Left, levelNode.level + 1})
		}

		if levelNode.node.Right != nil {
			queue.offer(&LevelNode{levelNode.node.Right, levelNode.level + 1})
		}
	}

	result = append(result, levelList)

	return result
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Println(zigzagLevelOrder(root))
}
