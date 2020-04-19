package main

import "fmt"

func longestPalindrome(s string) string {

	m := len(s)
	if m <= 1 {
		return s
	}

	//configure dp mark palindrome or not
	dp := initDp(m)
	result := ""
	for i := 0; i < m; i++ {
		for j := i; j >= 0; j-- {
			if s[i] == s[j] {
				if i-j <= 2 || dp[j+1][i-1] {
					dp[j][i] = true
					if i-j+1 > len(result) {
						fmt.Println(i, j)
						result = s[j : i+1]
					}
				}
			}
		}
	}

	fmt.Println(dp)

	return result
}

func initDp(m int) [][]bool {
	dp := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, m)
	}

	return dp
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("abcda"))
}
