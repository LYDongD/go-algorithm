package main

import "fmt"

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	lineDic, columnDic := make(map[int]int), make(map[int]int)

	//mark line and column to clear
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				lineDic[i]++
				columnDic[j]++
			}
		}
	}

	//clear by line
	for i := range lineDic {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = 0
		}
	}

	//clear by column
	for j := range columnDic {
		for i := 0; i < len(matrix); i++ {
			matrix[i][j] = 0
		}
	}
}

func main() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
