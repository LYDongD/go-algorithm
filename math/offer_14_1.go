package main

import "fmt"

func cuttingRope(n int) int {

	max := 0
	//cut from 2 -> n segments
	for i := 2; i <= n; i++ {
		s := maxCount(n, i)
		if s > max {
			max = s
		}
	}
	return max
}

func maxCount(n, m int) int {
	if m == 1 {
		return n
	}
	return (n / m) * maxCount(n-n/m, m-1)
}

func main() {
	fmt.Println(cuttingRope(2))
	fmt.Println(cuttingRope(10))
	fmt.Println(cuttingRope(8))
}
