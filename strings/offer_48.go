package main

import "fmt"
import "strings"

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	currentSubstr := ""
	maxLength := 0

	for i := 0; i < len(s); i++ {
		char := s[i : i+1]
		//check repeat
		if index := strings.Index(currentSubstr, char); index >= 0 {
			currentSubstr = currentSubstr[index+1:] + char
		} else {
			currentSubstr = currentSubstr + char
			if len(currentSubstr) > maxLength {
				maxLength = len(currentSubstr)
			}
		}
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring("aaaaaa"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
