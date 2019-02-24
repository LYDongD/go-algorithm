package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows < 2 {
		panic("at least 3")
	}

	array := make([][]rune, numRows)
	index := 0
	endFlag := false
	for !endFlag {
		for j := 0; j < numRows-2; j++ {
			array = addOneColumn(array)
			if j == 0 {
				for i := 0; i < numRows; i++ {
					array[i][j] = rune(s[index])
					index++
					if index == len(s) {
						endFlag = true
						break
					}
				}
			}

			fmt.Println(array, numRows-1-j, j, index, len(s))
			array[numRows-1-j][j] = rune(s[index])
			index++
			if index == len(s) {
				endFlag = true
				break
			}

		}
	}

	result := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(array[0]); j++ {
			if array[i][j] != '!' {
				result = result + string(array[i][j])
			}
		}
	}

	return result
}

func addOneColumn(array [][]rune) [][]rune {
	for index, rowArr := range array {
		if len(rowArr) == 0 {
			rowArr = []rune{'!'}
		} else {
			rowArr = append(rowArr, '!')
		}
		array[index] = rowArr
	}

	return array
}

func main() {
	fmt.Println(convert("PAHNAPLSIIGYIR", 3))
}
