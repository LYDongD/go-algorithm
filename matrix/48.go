package main

import "fmt"

func rotate(matrix [][]int)  {
    if len(matrix) == 0 || len(matrix) != len(matrix[0]) {
	return
    }

    N := len(matrix)
    for i := 0; i < N / 2; i++ {
	for j := i; j < N - i - 1; j++ {
	    r, c := i, j
	    last := matrix[r][c]
	    for k := 0; k < 4; k++ {
		t := c
		c = N - r - 1
		r = t
		next := matrix[r][c]
		matrix[r][c] = last
		last = next
	    }
	}
    }
}

func main() {
    matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}
    rotate(matrix)
    fmt.Println(matrix)


    matrix2 := [][]int{{5, 1, 9,11},{2, 4, 8,10},{13, 3, 6, 7}, {15,14,12,16}}
    rotate(matrix2)
    fmt.Println(matrix2)

}
