package main

import "fmt"

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}

	max := 1
	for n > 4 {
		max = max * 3 % 1000000007
		n = n - 3
	}

	return max * n % 1000000007

}

func main() {
	fmt.Println(cuttingRope(2))
	fmt.Println(cuttingRope(10))
	fmt.Println(cuttingRope(120))
}
