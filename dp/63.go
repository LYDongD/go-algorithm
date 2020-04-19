package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
	return 0
    }

    m, n := len(obstacleGrid), len(obstacleGrid[0])
    sum := initSumMatrix(m, n)
    for i := m - 1; i >= 0; i-- {
	for j := n - 1; j >= 0; j-- {
	    if obstacleGrid[i][j] == 1 {
		sum[i][j] = 0
		continue
	    }

	    if i == m - 1 && j == n - 1 {
		sum[i][j] = 1
	    }else {
		sum[i][j] = sum[i+1][j] + sum[i][j+1]
	    }

	}
    }

    return sum[0][0]
}

func initSumMatrix(m, n int) [][]int {
    sum := make([][]int, m + 1)
    for i := 0; i <= m; i++ {
	sum[i] = make([]int, n + 1)
    }

    return sum
}

func main() {
   fmt.Println( uniquePathsWithObstacles([][]int{{0,0,0},{0,1,0},{0,0,0}}))
   fmt.Println( uniquePathsWithObstacles([][]int{{1}}))
}
