package main

import "fmt"

func verifyPostorder(postorder []int) bool {
	if len(postorder) == 0 {
		return true
	}

	return check(postorder, 0, len(postorder)-1)
}

func check(postorder []int, startIndex, rootIndex int) bool {
	if startIndex >= rootIndex {
		return true
	}

	root := postorder[rootIndex]

	//find divide line
	divideIndex := rootIndex
	for i := startIndex; i < rootIndex; i++ {
		if postorder[i] > root {
			divideIndex = i
			break
		}
	}

	//check
	for i := divideIndex; i < rootIndex; i++ {
		if postorder[i] < root {
			return false
		}
	}

	//check left and right tree recursively
	return check(postorder, startIndex, divideIndex-1) && check(postorder, divideIndex, rootIndex-1)
}

func main() {
	fmt.Println(verifyPostorder([]int{1, 6, 3, 2, 5}))
	fmt.Println(verifyPostorder([]int{1, 3, 2, 6, 5}))
	fmt.Println(verifyPostorder([]int{1, 2, 5, 10, 6, 9, 4, 3}))
}
