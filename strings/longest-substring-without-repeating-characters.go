package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
    n := len(s)
    if n <= 1 {
	return n
    }

    maxCount := 0
    foundMap := make(map[rune]int)
    start := 0
    for index, char := range s {
	prevIndex, ok := foundMap[char]
	if !ok {
	    foundMap[char] = index
	}else {
	    maxCount = max(maxCount, index - start)
	    start = max(start, prevIndex + 1)
	    foundMap[char] = index
	}
    }

    maxCount = max(maxCount, n - start)

    return maxCount
}

func max (a, b int) int {
    if a > b {
	return a
    }

    return b
}

func main() {
    fmt.Println(lengthOfLongestSubstring("abcabcbb"))
    fmt.Println(lengthOfLongestSubstring("bbbbb"))
    fmt.Println(lengthOfLongestSubstring("pwwke"))
}


