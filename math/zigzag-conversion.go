package main

import "fmt"

func convert(s string, numRows int) string {
    if numRows < 2 {
	return s
    }

    array := make([][]rune, numRows)
    index := 0
    endFlag := false
    run := 0
    for !endFlag {
	for j := 0; j < numRows-1 && !endFlag; j++ {
	    array = addOneColumn(array)
	    if j == 0 {
		for i := 0; i < numRows; i++ {
		    array[i][j + run * (numRows-1) ]= rune(s[index])
		    index++
		    if index == len(s) {
			endFlag = true
			break
		    }
		}
	    }else {
		array[numRows-1-j][j + run * (numRows - 1) ] = rune(s[index])
		index++
		if index == len(s) {
			endFlag = true
			break
		}
	    }

	}
	run++
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
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
}
