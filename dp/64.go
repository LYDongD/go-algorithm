package main

import "fmt"

func minPathSum(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
	return 0
    }

    m, n := len(grid), len(grid[0])
    sum := initSum(m,n)
    for i := m - 1; i >= 0; i-- {
	for j := n - 1; j >= 0; j-- {
	    if i == m - 1 && j == n -1 {
		sum[i][j] = grid[i][j]
	    }else if i == m -1 {
		sum[i][j] = grid[i][j] + sum[i][j+1]
	    }else if j == n -1 {
		sum[i][j] = grid[i][j] + sum[i+1][j]
	    }else {
		sum[i][j] = min(grid[i][j] + sum[i + 1][j], grid[i][j] + sum[i][j + 1])
	    }
	    fmt.Println(i,j,sum[i][j])
	}
    }

    return sum[0][0]
}

func initSum(m, n int) [][]int{
    sum := make([][]int, m)
    for i := 0; i < len(sum); i++ {
	sum[i] = make([]int, n)
    }

    return sum
}

func min(a, b int) int {
    if a < b {
	return a
    }

    return b
}

func main() {
    fmt.Println(minPathSum([][]int{{1,3,1},{1,5,1},{4,2,1}}))
}
