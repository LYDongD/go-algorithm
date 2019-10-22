package main

import "fmt"

func bulbSwitch(n int) int {
	return sqrt(n)
}

func sqrt(num int) int {
	if num == 1 {
		return 1
	}

	start, end := 1, num
	for start <= end {
		mid := start + (end-start)/2
		if mid*mid == num {
			return mid
		}

		if mid*mid < num {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	if start*start < num {
		return start
	}

	return start - 1
}

func main() {
	fmt.Println(sqrt(14))
	fmt.Println(sqrt(15))
	fmt.Println(sqrt(16))
	fmt.Println(sqrt(17))
	fmt.Println(sqrt(18))
	fmt.Println(sqrt(19))
	fmt.Println(sqrt(20))
	fmt.Println(sqrt(21))
	fmt.Println(sqrt(22))
	fmt.Println(sqrt(23))
	fmt.Println(sqrt(24))
	fmt.Println(sqrt(25))
}
