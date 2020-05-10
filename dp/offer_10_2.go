package main

import "fmt"

func numWays(n int) int {

	if n == 0 {
		return 1
	}

	if n <= 2 {
		return n
	}

	pre, current := 1, 2
	sum := 0
	for i := 3; i <= n; i++ {
		sum = (pre + current) % 1000000007
		pre = current
		current = sum
	}

	return sum
}

func main() {
	fmt.Println(numWays(2))
	fmt.Println(numWays(7))
}
