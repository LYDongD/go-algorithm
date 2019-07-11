package main

import "fmt"

func singleNumber(nums []int) int {
	a, b := 0, 0
	for i := 0; i < len(nums); i++ {
		b = (b ^ nums[i]) & ^a
		a = (a ^ nums[i]) & ^b
		fmt.Println(b, a)
	}
	return b
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
	fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
	fmt.Println(singleNumber([]int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}))
}
