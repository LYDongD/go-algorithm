package main

import "fmt"

func minCostClimbingStairs(cost []int) int {

	n := len(cost)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1]+cost[i], dp[i-2]+cost[i])
	}

	return min(dp[n-2], dp[n-1])
}

func min(a int, b int) int {
	if a <= b {
		return a
	}

	return b
}

func main() {

	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))

}
