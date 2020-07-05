package main

import "fmt"

func nthUglyNumber(n int) int {
	if n <= 0 {
		panic("n must be positive")
	}

	count := 0
	uglys := []int{1, 2, 3, 4, 5, 6}
	cursor2, cursor3, cursor5 := 3, 2, 4
	if n <= len(uglys) {
		return uglys[n-1]
	}

	count = 6

	//generate by 2,3 and 5
	i := 5
	for count < n {
		i++
		result2 := 2 * uglys[cursor2]
		result3 := 3 * uglys[cursor3]
		result5 := 5 * uglys[cursor5]

		minResult := min(result2, result3, result5)
		uglys = append(uglys, minResult)
		if minResult == result2 {
			cursor2++
		}

		if minResult == result3 {
			cursor3++
		}

		if minResult == result5 {
			cursor5++
		}

		count++
	}

	return uglys[len(uglys)-1]
}

func min(a, b, c int) int {
	min := a
	if b < min {
		min = b
	}

	if c < min {
		min = c
	}

	return min
}

func main() {
	fmt.Println(nthUglyNumber(10))
	fmt.Println(nthUglyNumber(1352))
}
