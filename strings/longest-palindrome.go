package main

import "fmt"

func longestPalindrome(s string) int {
    if len(s) <= 1 {
	return len(s)
    }

    //char count
    charCountMap := make(map[rune]int)
    for _, char := range s {
	charCountMap[char]++
    }

    //find odd group and add up
    result := 0
    singleFlag := false
    for _, count := range charCountMap {
	if count % 2 == 0 {
	    result += count
	}else {
	    result += count / 2 * 2
	    singleFlag = true
	}
    }

    //add one single middle 
    if singleFlag {
	result++
    }

    return result
}

func main() {
    fmt.Println(longestPalindrome("abccccdd"))
}
