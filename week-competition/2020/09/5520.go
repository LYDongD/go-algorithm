package main

import "fmt"

func maxUniqueSplit(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	candidateDic := make(map[string]int)
	return findMax(s, 0, 0, candidateDic)
}

func findMax(s string, start, count int, candidateDic map[string]int) int {
	if start >= len(s) {
		return count
	}

	maxCount := 0
	for i := start + 1; i <= len(s); i++ {
		sub := s[start:i]
		if candidateDic[sub] == 0 {
			candidateDic[sub] = 1
			maxCount = max(maxCount, findMax(s, i, count+1, candidateDic))
			candidateDic[sub] = 0
		}
	}

	return maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(maxUniqueSplit("ababccc"))
}
