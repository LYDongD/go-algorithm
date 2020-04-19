package main

import "fmt"

func uniquePaths(m int, n int) int {
	row, col := n-1, m-1
	sum := make([][]int, n+1)
	for i := 0; i < len(sum); i++ {
		sum[i] = make([]int, m+1)
	}

	for i := row; i >= 0; i-- {
		for j := col; j >= 0; j-- {
			if i == row && j == col {
				sum[i][j] = 1
			} else {
				sum[i][j] = sum[i+1][j] + sum[i][j+1]
			}
		}
	}

	return sum[0][0]
}

func main() {
	fmt.Println(uniquePaths(3, 2))
}
