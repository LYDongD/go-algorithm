package main

import "fmt"

func minDays(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 || n == 3 {
		return 2
	}

	minDays := make([]int64, 3)
	minDays[0], minDays[1], minDays[2] = 1, 2, 2
	for i := int64(3); i <= n-1; i++ {
		m := i + 1
		if m%3 == 0 && m%2 == 0 {
			minDays[i] = minIn3(minDays[m-1-1], minDays[m/2-1], minDays[m-2*(m/3)-1]) + 1
		} else if m%2 == 0 {
			minDays[i] = minIn2(minDays[m-1-1], minDays[m/2-1]) + 1
		} else if m%3 == 0 {
			minDays[i] = minIn2(minDays[m-1-1], minDays[m-2*(m/3)-1]) + 1
		} else {
			minDays[i] = minDays[m-1-1] + 1
		}
	}

	return int(minDays[n-1])
}

func minIn3(a, b, c int64) int64 {
	min := a
	if min > b {
		min = b
	}

	if min > c {
		min = c
	}

	return min
}

func minIn2(a, b int64) int64 {
	if a < b {
		return a
	}

	return b
}

func main() {
	//fmt.Println(minDays(6))
	//fmt.Println(minDays(9))
	//fmt.Println(minDays(10))
	//fmt.Println(minDays(6))
	//fmt.Println(minDays(1))
	fmt.Println(minDays(84806671))
}
