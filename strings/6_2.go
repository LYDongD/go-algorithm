package main

import "fmt"

func convert(s string, numRows int) string {
	if len(s) <= 1 || numRows <= 1 {
		return s
	}

	n := min(len(s), numRows)
	strArr := make([]string, n)
	row, up := 0, false
	for i := 0; i < len(s); i++ {
		if row == 0 || row == numRows-1 {
			up = !up
		}

		strArr[row] = strArr[row] + s[i:i+1]

		if up {
			row++
		} else {
			row--
		}
	}

	result := ""
	for _, str := range strArr {
		result = result + str
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
}
