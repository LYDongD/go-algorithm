package main

import "fmt"

func invoke(sum int) int {
	if sum == 0 {
		return 0
	}
	return invoke(sum-1) + 1
}

var a int = 0

func invoke2(sum int) {
	if sum == 0 {
		return
	}

	a++
	invoke2(sum - 1)
}

func main() {
	fmt.Println(invoke(10))
	fmt.Println(invoke(10))
}
