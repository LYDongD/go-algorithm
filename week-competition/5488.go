package main

import "fmt"

func minOperations(n int) int {
	target := n
	count := 0
	for i := 0; i <= n/2-1; i++ {
		count += target - (2*i + 1)
	}

	return count
}

func main() {
	fmt.Println(minOperations(3))
	fmt.Println(minOperations(6))
}
