package main

import "fmt"

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value int
	count int
}

func findAdjacentBiggerVal(num, adjacentNum int, node *TreeNode) int {

	if node == nil {
		return adjacentNum
	}

	if node.value == num {
		return node.value
	}

	if node.value > num {
		adjacentNum = node.value
		return findAdjacentBiggerVal(num, adjacentNum, node.left)
	}

	return findAdjacentBiggerVal(num, adjacentNum, node.right)
}

func findAdjacentSmallerVal(num, adjacentNum int, node *TreeNode) int {
	if node == nil {
		return adjacentNum
	}

	if node.value == num {
		return adjacentNum
	}

	if node.value < num {
		adjacentNum = node.value
		return findAdjacentSmallerVal(num, adjacentNum, node.right)
	}

	return findAdjacentSmallerVal(num, adjacentNum, node.left)
}

func insert(num int, root *TreeNode) *TreeNode {
	if root == nil {
		return &TreeNode{value: num, count: 1}
	}

	if root.value == num {
		root.count++
	} else if root.value > num {
		root.left = insert(num, root.left)
	} else {
		root.right = insert(num, root.right)
	}

	return root
}

func remove(num int, root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root.value == num {
		root.count--
		if root.count == 0 {
			if root.left == nil && root.right == nil {
				root = nil
			} else if root.left == nil {
				root = root.right
			} else if root.right == nil {
				root = root.left
			} else {
				node := removeLeftSmallest(root.left)
				node.left = root.left
				node.right = root.right
				root = node
			}
		}
		return root
	}

	if root.value > num {
		root.left = remove(num, root.left)
	} else {
		root.right = remove(num, root.right)
	}

	return root
}

func removeLeftSmallest(root *TreeNode) *TreeNode {
	if root.left == nil {
		return root
	}

	node := root
	for {
		if node.left.left == nil {
			leftSmallest := node.left
			node.left = nil
			return leftSmallest
		}

		node = node.left
	}
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) <= 1 || k <= 0 || t < 0 {
		return false
	}

	var root *TreeNode
	for i, num := range nums {
		smaller := findAdjacentSmallerVal(num, num+1, root)
		if smaller <= num && abs(num-smaller) <= t {
			return true
		}
		bigger := findAdjacentBiggerVal(num, num-1, root)
		if bigger >= num && abs(num-bigger) <= t {
			return true
		}

		if i >= k {
			root = remove(nums[i-k], root)
		}

		root = insert(num, root)
	}

	return false
}

func abs(a int) int {
	if a >= 0 {
		return a
	}

	return -a
}

func main() {

	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 0, 1, 1}, 1, 2))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))

}
