package main

import "fmt"

func rotatedDigits(N int) int {
	if N < 1 {
		return 0
	}

	count := 0
	for num := 1; num <= N; num++ {
		if isGood(num) {
			count++
		}
	}

	return count
}

func isGood(num int) bool {

	rotateDifferFlag := false
	for num > 0 {
		digit := num % 10
		if digit == 1 || digit == 0 || digit == 8 {
			num = num / 10
			continue
		}

		if digit == 2 || digit == 5 || digit == 6 || digit == 9 {
			num = num / 10
			rotateDifferFlag = true
			continue
		}

		return false
	}

	return rotateDifferFlag
}

func main() {
	//fmt.Println(rotatedDigits(10))
	fmt.Println(rotatedDigits(857))
}
