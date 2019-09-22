package main

import "fmt"

func nthUglyNumber(n int) int {
	if n <= 6 {
		return n
	}

	ugly := make([]int, n)
	ugly[0] = 1
	factor1, factor2, factor3 := 2, 3, 5
	index1, index2, index3 := 0, 0, 0
	for i := 1; i < n; i++ {
		num := min(factor1, factor2, factor3)
		ugly[i] = num
		if num == factor1 {
			index1++
			factor1 = 2 * ugly[index1]
		}

		if num == factor2 {
			index2++
			factor2 = 3 * ugly[index2]
		}

		if num == factor3 {
			index3++
			factor3 = 5 * ugly[index3]
		}
	}

	return ugly[n-1]

}

func min(a, b, c int) int {
	minVal := a
	if b < a {
		minVal = b
	}

	if c < minVal {
		minVal = c
	}

	return minVal
}

func main() {
	fmt.Println(nthUglyNumber(10))
}
