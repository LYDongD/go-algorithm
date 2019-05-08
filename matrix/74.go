package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])

	//find goal row
	start, end := 0, m-1
	for start <= end {
		mid := start + (end-start)/2

		if matrix[mid][0] <= target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	if start == 0 {
		return false
	}

	goalRow := start - 1

	//find goal
	start, end = 0, n-1
	for start <= end {
		mid := start + (end-start)/2
		if matrix[goalRow][mid] == target {
			return true
		}

		if matrix[goalRow][mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1}, {3}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 1}}, 0))
	fmt.Println(searchMatrix([][]int{{1}}, 1))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 13))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 25))
}
