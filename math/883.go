package  main

import "fmt"

func projectionArea(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0{
	return 0
    }

    sumXY := 0
    sumYZ := 0
    //遍历行
    for i := 0; i < len(grid); i++ {
	//遍历列, 取出当前行的最大高度作为YZ的高度
	maxZ := 0
	for j := 0; j < len(grid[i]); j++ {
	    if grid[i][j] > 0 {
		sumXY++
	    }
	    if grid[i][j] > maxZ {
		maxZ = grid[i][j]
	    }
	}
	sumYZ += maxZ
    }

    //遍历列
    sumXZ := 0
    for k := 0; k < len(grid[0]); k++ {
	maxZ := 0
	for t := 0; t < len(grid); t++ {
	    if grid[t][k] > maxZ {
		maxZ = grid[t][k]
	    }
	}
	sumXZ += maxZ
    }

    return sumXY + sumYZ + sumXZ
}

func main() {
    fmt.Println(projectionArea([][]int{{2}}))
    fmt.Println(projectionArea([][]int{{1,2}, {3,4}}))
    fmt.Println(projectionArea([][]int{{1,0}, {0,2}}))
    fmt.Println(projectionArea([][]int{{1,1,1}, {1,0,1}, {1,1,1}}))
    fmt.Println(projectionArea([][]int{{2,2,2}, {2,1,2}, {2,2,2}}))
}
