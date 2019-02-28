package main

import "fmt"
import "math"

func largestTriangleArea(points [][]int) float64 {
    if len(points) == 0 {
	return float64(0)
    }

    max := float64(0)
    for i := 0; i < len(points); i++ {
	for j := i + 1; j < len(points); j++ {
	    for k := j + 1; k < len(points); k++ {
		s := square(i, j, k, points)
		if s > max {
		    max = s
		}
	    }
	}
    }

    return max
}

func square(i, j, k int, points [][]int) float64 {
    return math.Abs(float64((points[i][0] - points[k][0]) * (points[j][1] - points[k][1]) -
    (points[j][0] - points[k][0]) * (points[i][1] - points[k][1]))) / 2.0
}

func main() {
    fmt.Println(largestTriangleArea([][]int{{0,0},{0,1},{1,0},{0,2},{2,0}}))
}
