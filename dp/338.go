package main

import "fmt"

func countBits(num int) []int {
	dp := make([]int, num+1)
	if num < 0 {
		return dp
	}

	if num == 0 {
		return []int{0}
	}

	if num == 1 {
		return []int{0, 1}
	}

	dp[0], dp[1] = 0, 1

	base := 1
	for i := 2; i <= num; i++ {
		if i == base*2 {
			dp[i] = 1
			base = i
		} else {
			dp[i] = dp[i-base] + 1
		}

	}

	return dp
}

func main() {
	//fmt.Println(countBits(2))
	//fmt.Println(countBits(5))
	fmt.Println(countBits(8))
	//fmt.Println(countBits(11))
	//fmt.Println(countBits(12))
	//fmt.Println(countBits(13))
}
