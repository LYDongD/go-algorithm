package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	maxEdge := 1
	found := false
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '0' {
				continue
			}

			found = true
			for k := 1; k <= rows-i && k <= cols-j; k++ {
				lastRow := i + k - 1
				lastCol := j + k - 1
				shouldBreak := false
				for t := j; t <= lastCol; t++ {
					if matrix[lastRow][t] == '0' {
						shouldBreak = true
						break
					}
				}

				if shouldBreak {
					break
				}

				for t := i; t <= lastRow; t++ {
					if matrix[t][lastCol] == '0' {
						shouldBreak = true
						break
					}
				}

				if shouldBreak {
					break
				}

				if k > maxEdge {
					maxEdge = k
				}
			}
		}
	}

	if found {
		return maxEdge * maxEdge
	}

	return 0
}

func main() {
	fmt.Println(maximalSquare([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
	fmt.Println(maximalSquare([][]byte{{'1', '1', '0', '1'}, {'1', '1', '0', '1'}, {'1', '1', '1', '1'}}))
}
