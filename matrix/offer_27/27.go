package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return result
	}

	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for {
		//top: left -> right
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}

		//right: top -> bottom
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if right < left {
			break
		}

		//bottom: right -> left
		for i := right; i >= left; i-- {
			result = append(result, matrix[bottom][i])
		}
		bottom--
		if bottom < top {
			break
		}

		//left: bottom -> top
		for i := bottom; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return result
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(matrix))
	matrix2 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(matrix2))
}
