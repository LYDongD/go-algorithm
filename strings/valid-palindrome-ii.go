package main

import "fmt"

func validPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	start := 0
	end := len(s) - 1
	for start < end {
		if s[start] == s[end] {
			start++
			end--
			continue
		}

		return validPalindromeWithoutDelete(s, start+1, end) || validPalindromeWithoutDelete(s, start, end-1)
	}

	return true
}

func validPalindromeWithoutDelete(s string, start int, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}

	return true
}

func main() {

	fmt.Println(validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"))
}
