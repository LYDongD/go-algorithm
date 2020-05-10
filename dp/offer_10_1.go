package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}

	pre, current := 0, 1
	sum := 0
	for i := 2; i <= n; i++ {
		sum = (pre + current) % 1000000007
		pre = current
		current = sum
	}

	return sum
}

func main() {
	fmt.Println(fib(5))
	fmt.Println(fib(2))
}
