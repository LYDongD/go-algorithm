package main

import "fmt"

func longestPalindrome(s string) int {
	if s == "" {
		return 0
	}

	countMap := make(map[byte]int)
	for _, char := range s {
		countMap[byte(char)]++
	}

	sum := 0
	for char := range countMap {
		if countMap[char]%2 == 0 {
			sum += countMap[char]
		} else {
			sum += countMap[char] - 1
		}
	}

	if sum == len(s) {
		return sum
	} else {
		return sum + 1
	}
}

// func main() {
// 	fmt.Println(longestPalindrome("aba"))
// }
