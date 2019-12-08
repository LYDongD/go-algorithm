package main

import "fmt"

func longestPalindrome(s string) string {

	n := len(s)

	if n < 2 {
		return s
	}

	//iterate all substring from short to long
	palindrome := make([][]bool, n)
	for i := 0; i < len(palindrome); i++ {
		palindrome[i] = make([]bool, n)
	}

	result := ""
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 2 || palindrome[i+1][j-1] {
					palindrome[i][j] = true
					if j-i+1 > len(result) {
						result = s[i : j+1]
					}
				}
			}
		}
	}

	return result
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("ac"))
	fmt.Println(longestPalindrome("aaaa"))
	fmt.Println(longestPalindrome("abcba"))
}
