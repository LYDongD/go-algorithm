package main

import (
	"fmt"
	"unicode"
)

func reverseWords(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	start := -1
	length := 0
	chars := []rune(s)
	for index, char := range chars {
		if unicode.IsSpace(char) && length > 0 {
			reverse(chars, start, length)
			length = 0
			start = -1
		} else if !unicode.IsSpace(char) {
			length++
			if start < 0 {
				start = index
			}
		}
	}

	if length > 0 {
		reverse(chars, start, length)
	}

	return string(chars)
}

func reverse(chars []rune, start int, length int) {
	end := start + length - 1
	for start < end {
		chars[start], chars[end] = chars[end], chars[start]
		start++
		end--
	}
}

func main() {

	fmt.Println(reverseWords("Let's take LeetCode contest"))

}
