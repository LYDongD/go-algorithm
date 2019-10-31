package main

import "fmt"

func isValidSerialization(preorder string) bool {
	if len(preorder) == 0 {
		return false
	}

	if isNull(preorder[0]) {
		return len(preorder) == 1
	}

	diff := 1
	for i := 0; i < len(preorder); i++ {
		node := preorder[i]
		if isComma(node) || isNotComplete(node, i, preorder) {
			continue
		}

		diff--
		if diff < 0 {
			return false
		}

		if !isNull(node) {
			diff += 2
		}
	}

	return diff == 0
}

func isNull(node byte) bool {
	return node == '#'
}

func isComma(node byte) bool {
	return node == ','
}

func isNotComplete(node byte, i int, preorder string) bool {
	if i == len(preorder)-1 {
		return false
	}

	if isComma(preorder[i+1]) {
		return false
	}

	return true
}

func main() {
	fmt.Println("vim-go")
	fmt.Println(isValidSerialization("9,#,92,#,#"))
	fmt.Println(isValidSerialization("#"))
	fmt.Println(isValidSerialization("#,#,#"))
	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
	fmt.Println(isValidSerialization("1,#"))
	fmt.Println(isValidSerialization("9,#,#,1"))
}
