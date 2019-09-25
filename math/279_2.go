package main

import "fmt"

func numSquares(n int) int {
	if n <= 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		min := int(^uint(0) >> 1)
		for j := 1; j*j <= i; j++ {
			if dp[i-j*j]+1 < min {
				min = dp[i-j*j] + 1
			}
		}
		dp[i] = min
	}

	return dp[n]
}

func main() {
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
}
