package main

import "fmt"

func minMoves(nums []int) int {

	sum, min := findSumAndMin(nums)
	return sum - min*len(nums)
}

func findSumAndMin(nums []int) (sum int, min int) {

	sum = 0
	min = nums[0]
	for _, num := range nums {
		sum += num
		if num < min {
			min = num
		}
	}

	return sum, min
}

func main() {

	nums := []int{1, 2, 3}
	fmt.Println(findSumAndMin(nums))
	fmt.Println(minMoves(nums))
}
