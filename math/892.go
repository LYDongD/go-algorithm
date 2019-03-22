package main

import "fmt"

func surfaceArea(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
	return 0
    }

    sum := 0
    for i := 0; i < len(grid); i++ {
	for j := 0; j < len(grid[i]); j++ {
	    if grid[i][j] > 0 {
		sum += 2
		//front
		if i - 1 >= 0 {
		    sum += max(0, grid[i][j] - grid[i-1][j])
		}else {
		    sum += grid[i][j]
		}
		//back
		if i + 1 < len(grid) {
		    sum += max(0, grid[i][j] - grid[i+1][j])
		}else {
		    sum += grid[i][j]
		}
		//left
		if j - 1 >= 0 {
		    sum += max(0, grid[i][j] - grid[i][j-1])
		}else {
		    sum += grid[i][j]
		}

		//right
		if j + 1 < len(grid[i]) {
		    sum += max(0, grid[i][j] - grid[i][j+1])
		}else {
		    sum += grid[i][j]
		}
	    }
	}
    }

    return sum
}

func max(a, b int) int {
    if a > b {
	return a
    }

    return b
}


func main() {
    fmt.Println(surfaceArea([][]int{{2}}))
    fmt.Println(surfaceArea([][]int{{1,2}, {3,4}}))
    fmt.Println(surfaceArea([][]int{{1,1,1}, {1,0,1}, {1,1,1}}))
    fmt.Println(surfaceArea([][]int{{2,2,2}, {2,1,2}, {2,2,2}}))
}
