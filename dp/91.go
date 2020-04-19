package main

import "fmt"

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	n := len(s)
	sum := make([]int, n+1)
	sum[0], sum[1] = 1, 1
	for i := 2; i <= n; i++ {
		single, double := 0, 0
		if s[i-1:i] > "0" {
			single = sum[i-1]
		}

		if s[i-2:i] <= "26" && s[i-2:i] >= "10" {
			double = sum[i-2]
		}

		sum[i] = single + double
	}

	return sum[n]
}

func main() {
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("00"))
	fmt.Println(numDecodings("01"))
}
