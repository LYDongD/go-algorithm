package main

import "fmt"

var minSum int

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}

	//dp from bottom to up by layer, left -> right at layer
	dp := make([]int, len(triangle))
	lastRow := len(triangle) - 1
	for i := 0; i < len(triangle[lastRow]); i++ {
		dp[i] = triangle[lastRow][i]
	}

	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col < len(triangle[row]); col++ {
			dp[col] = min(dp[col], dp[col+1]) + triangle[row][col]
		}
	}

	return dp[0]

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}
