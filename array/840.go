package main

import "fmt"

func numMagicSquaresInside(grid [][]int) int {
    if len(grid) < 3 || len(grid[0]) < 3 {
	return 0
    }

    total := 0
    for i := 0; i <= len(grid) - 3; i++ {
	for j := 0; j <= len(grid[0]) - 3; j++ {
	    sum := grid[i][j] + grid[i+1][j+1] + grid[i+2][j+2]
	    if rowSumEqual(sum, i, j, grid) && colSumEqual(sum, i, j, grid) &&
	    grid[i+2][j] + grid[i+1][j+1] + grid[i][j+2] == sum {
		total++
	    }
	}
    }

    return total
}

func rowSumEqual(sum, i, j  int, grid [][]int) bool {
    num := grid[i][j]
    for k := i; k <= i + 2; k++ {
	rowSum := 0
	for t := j; t <= j + 2; t++ {
	    if (k != i && t != j &&  grid[k][t] == num) || 
	    grid[k][t] < 1 || grid[k][t] > 9 {
		return false
	    }
	    rowSum += grid[k][t]
	}

	if rowSum != sum {
	    return false
	}
    }

    return true
}

func colSumEqual(sum, i, j int, grid [][]int) bool {
    for k := j; k <= j + 2; k++ {
        colSum := 0
        for t := i; t <= i + 2; t++ {
            colSum += grid[t][k]
        }

        if colSum != sum {
            return false
        }
    }

    return true
}

func main() {
    fmt.Println(numMagicSquaresInside([][]int{{4,3,8,4},{9,5,1,9},{2,7,6,2}}))
}
