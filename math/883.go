package  main

import "fmt"

func projectionArea(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0{
	return 0
    }

    sum  := 0
    for i := 0; i < len(grid); i++ {
	for j := 0; j < len(grid[i]); j++ {
	    square := 1 + 1 * grid[i][j] + 1 * grid[i][j]
	    sum += square
	}
    }

    return sum
}

func main() {
    fmt.Println(projectionArea([][]int{{2}}))
    fmt.Println(projectionArea([][]int{{1,2}, {3,4}}))
}
