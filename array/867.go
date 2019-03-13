package main

import "fmt"

func transpose(A [][]int) [][]int {
    if len(A) == 0 {
	return [][]int{}
    }

    result := make([][]int, len(A[0]))
    for i := 0; i < len(result); i++ {
	result[i] = make([]int, len(A))
	for j := 0; j < len(A); j++ {
	    result[i][j] = A[j][i]
	}
    }

    return result
}

func main() {
    fmt.Println(transpose([][]int{{1,2,3},{4,5,6},{7,8,9}}))
    fmt.Println(transpose([][]int{{1,2,3},{4,5,6}}))
}
