package main

import "fmt"

func flipAndInvertImage(A [][]int) [][]int {
    result := [][]int{}
    if len(A) == 0 {
	return result
    }

    for _, row := range A {
	for i := 0; i <  (len(row) + 1) / 2; i++ {
	    j := len(row) - i - 1
	    row[i], row[j] = row[j] ^ 1, row[i] ^ 1
        }

	result = append(result, row)
    }

    return result
}

func main() {
    fmt.Println(flipAndInvertImage([][]int{{1,1,0}, {1,0,1}, {0,0,0}}))
}


