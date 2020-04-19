package main

import "fmt"

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}

	sum := make([]int, n+1)
	sum[0], sum[1] = 1, 1
	for i := 2; i < len(sum); i++ {
		sum[i] = sum[i-1] + sum[i-2]
	}

	return sum[n]
}

func main() {
	fmt.Println(climbStairs(3))
}
