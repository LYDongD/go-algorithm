package main

import "fmt"

func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) == 0 {
		return false
	}

	rows := len(matrix)
	columns := len(matrix[0])

	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			if matrix[i-1][j-1] != matrix[i][j] {
				return false
			}
		}
	}

	return true
}

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
	fmt.Println(isToeplitzMatrix(matrix))
	matrix2 := [][]int{{1, 2}, {1, 2}}
	fmt.Println(isToeplitzMatrix(matrix2))
}
