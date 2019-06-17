package main

import "fmt"

const INT_MAX = int(^uint(0) >> 1)

func minimumTotal(triangle [][]int) int {
    if len(triangle) == 0 || len(triangle[0]) == 0 {
	return 0
    }

    rowMin := INT_MAX
    pathSum := 0
    for r := 0; r < len(triangle); r++ {
	for c := 0; c < len(triangle[r]); c++ {
	    if triangle[r][c] < rowMin {
		rowMin = triangle[r][c]
	    }
	}

	if rowMin < INT_MAX {
	    pathSum += rowMin
	    rowMin = INT_MAX
	}
    }

    return pathSum
}

func main() {
    fmt.Println(minimumTotal([][]int{{2},{3,4},{6,5,7},{4,1,8,3}}))
}
