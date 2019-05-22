package main

import "fmt"

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	dp := make([]int, len(s)+1)

	//empty string -> 1 way to decode
	dp[0] = 1
	if s[0] == '0' {
		dp[1] = 0
	} else {
		dp[1] = 1
	}

	//handle from the 2nd char
	for i := 1; i < len(s); i++ {
		if !isSingleValid(s[i]) && !isValid(s[i-1], s[i]) {
			return 0
		}

		if isSingleValid(s[i]) {
			dp[i+1] += dp[i]
		}

		if isValid(s[i-1], s[i]) {
			dp[i+1] += dp[i-1]
		}
	}

	return dp[len(s)]
}

func isSingleValid(a byte) bool {
	return a >= '1' && a <= '9'
}

func isValid(a byte, b byte) bool {
	s := (a-'0')*10 + b - '0'
	return s >= 1 && s <= 26
}

func main() {
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("0"))
	fmt.Println(numDecodings("01"))
}
