package main

import "fmt"

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 || amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	dp := make([]int, amount)
	return findCount(dp, coins, amount)
}

func findCount(dp, coins []int, remind int) int {
	if remind < 0 {
		return -1
	}
	if remind == 0 {
		return 0
	}

	if dp[remind-1] > 0 {
		return dp[remind-1]
	}

	min := int(^uint(0) >> 1)
	for _, coin := range coins {
		lastCount := findCount(dp, coins, remind-coin)
		if lastCount >= 0 && min > lastCount {
			min = lastCount
		}
	}

	if min == int(^uint(0)>>1) {
		return -1
	}

	dp[remind-1] = min + 1
	return 1 + min
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
}
