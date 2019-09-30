package main

import "fmt"

type NumMatrix struct {
    recSum [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    numMatrix := NumMatrix{}

    if len(matrix) == 0 || len(matrix[0]) == 0 {
	return numMatrix
    }

    recSum := initRecSum(len(matrix), len(matrix[0]))
    //init recsum matrix
    for i := 1; i <= len(matrix); i++ {
	for j := 1; j <= len(matrix[0]); j++ {
	    recSum[i][j] = recSum[i-1][j] + recSum[i][j-1] - recSum[i-1][j-1] + matrix[i-1][j-1]
	}
    }

    numMatrix.recSum = recSum
    return numMatrix
}

func initRecSum(row, col int) [][]int {
    recSum := make([][]int, row+1)
    for i, _ := range recSum {
	recSum[i] = make([]int, col+1)
    }
    return recSum
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    return this.recSum[row2+1][col2+1] - this.recSum[row1][col2+1] - this.recSum[row2 + 1][col1] + this.recSum[row1][col1]
}

func main() {
    matrix := [][]int{{3,0,1,4,2},{5,6,3,2,1},{1,2,0,1,5},{4,1,0,1,7},{1,0,3,0,5}}
    obj := Constructor(matrix)
    fmt.Println(obj.SumRegion(2,1,4,3))
    fmt.Println(obj.SumRegion(1,1,2,2))
    fmt.Println(obj.SumRegion(1,2,2,4))
}
