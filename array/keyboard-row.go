package main

import (
	"fmt"
	"strings"
)

func findWords(words []string) []string {

	if len(words) == 0 {
		return nil
	}

	row1Str := "QWERTYUIOP"
	row2Str := "ASDFGHJKL"
	row3Str := "ZXCVBNM"

	rowMap := make(map[rune]int)
	configureRowMap(rowMap, row1Str, 1)
	configureRowMap(rowMap, row2Str, 2)
	configureRowMap(rowMap, row3Str, 3)

	fmt.Println(rowMap)

	var result []string
	for _, word := range words {
		lastRow := 0
		isSameRow := true
		for _, char := range strings.ToUpper(word) {
			row := rowMap[char]
			if lastRow == 0 {
				lastRow = row
			} else if lastRow != row {
				isSameRow = false
				break
			}
		}

		if isSameRow {
			result = append(result, word)
		}
	}

	return result

}

func configureRowMap(rowMap map[rune]int, rowStr string, row int) {
	for _, char := range rowStr {
		rowMap[char] = row
	}
}

func main() {

	words := []string{"Hello", "Alaska", "Dad", "Peace"}

	fmt.Println(findWords(words))
}
