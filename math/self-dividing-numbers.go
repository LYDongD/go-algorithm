package main

import "fmt"

func selfDividingNumbers(left int, right int) []int {
	if left > right {
		panic("illegal bounds")
	}

	var result []int
	for num := left; num <= right; num++ {
		if isSelfDivided(num) {
			result = append(result, num)
		}
	}

	return result
}

func isSelfDivided(num int) bool {
	a := num
	for a != 0 {
		bit := a % 10
		if bit == 0 || num%bit != 0 {
			return false
		}

		a = a / 10
	}

	return true
}

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}
