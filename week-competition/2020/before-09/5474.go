package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {
	if root == nil {
		return 0
	}

	count := 0
	leafDis(root, distance, &count)
	return count
}

func leafDis(root *TreeNode, distance int, count *int) ([]int, []int) {

	if root.Left == nil && root.Right == nil {
		return []int{0}, []int{}
	}

	leftDis, rightDis := []int{}, []int{}

	if root.Left != nil {
		left, right := leafDis(root.Left, distance, count)
		leftDis = addAndAppend(leftDis, left)
		leftDis = addAndAppend(leftDis, right)
	}

	if root.Right != nil {
		left2, right2 := leafDis(root.Right, distance, count)
		rightDis = addAndAppend(rightDis, left2)
		rightDis = addAndAppend(rightDis, right2)
	}

	for i := 0; i < len(leftDis); i++ {
		for j := 0; j < len(rightDis); j++ {
			if leftDis[i]+rightDis[j] <= distance {
				*count = *count + 1
			}
		}
	}

	return leftDis, rightDis
}

func addAndAppend(source []int, target []int) []int {

	for _, num := range target {
		source = append(source, num+1)
	}

	return source
}

func main() {
	fmt.Println("hi")
}
