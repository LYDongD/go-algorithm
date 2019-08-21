package main

import "fmt"

func rangeBitwiseAnd(m int, n int) int {
	if m == 0 {
		return 0
	}

	count := 0
	for m != n {
		m = m >> 1
		n = n >> 1
		count++
	}

    return m << uint(count)
}

func main() {
	fmt.Println(rangeBitwiseAnd(20000, 2147483647))
	fmt.Println(rangeBitwiseAnd(2, 4))
	fmt.Println(rangeBitwiseAnd(3, 3)
}
