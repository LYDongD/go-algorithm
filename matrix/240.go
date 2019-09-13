package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	t, maxRow, maxCol := 0, len(matrix)-1, len(matrix[0])-1
	for t <= maxRow && t <= maxCol {
		if matrix[t][t] > target {
			return false
		}

		k, goal := searchRow(matrix[t], t, maxCol, target)
		if target == goal {
			return true
		}
		maxCol = k

		k, goal = searchCol(matrix, t, maxRow, target)
		if target == goal {
			return true
		}

		maxRow = k
		t++
	}

	return false
}

func searchRow(nums []int, start, end, target int) (int, int) {

	for start < end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid, target
		}

		if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	if nums[start] < target {
		return start, nums[start]
	}

	return start - 1, nums[start-1]

}

func searchCol(matrix [][]int, start, end, target int) (int, int) {
	k := start
	for start < end {
		mid := start + (end-start)/2
		if matrix[mid][k] == target {
			return mid, target
		}

		if matrix[mid][k] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	if matrix[start][k] <= target {
		return start, matrix[start][k]
	}

	return start - 1, matrix[start][k]
}

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	fmt.Println(searchMatrix(matrix, 5))
	fmt.Println(searchMatrix(matrix, 20))
}
