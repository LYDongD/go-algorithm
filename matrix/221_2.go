package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	//init dp
	dp := make([][]int, rows+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, cols+1)
	}

	//generate dp
	maxEdge := 0
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if maxEdge < dp[i][j] {
					maxEdge = dp[i][j]
				}
			}
		}
	}

	return maxEdge * maxEdge
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	fmt.Println(maximalSquare([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
	fmt.Println(maximalSquare([][]byte{{'1', '1', '0', '1'}, {'1', '1', '0', '1'}, {'1', '1', '1', '1'}}))
}
