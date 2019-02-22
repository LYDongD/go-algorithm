package main

import "fmt"

func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 {
		return 0
	}

	//cal sum
	sum := 0
	for _, num := range nums {
		sum += num
	}

	currentSum := 0
	for index, num := range nums {
		currentSum += num
		leftSum := currentSum - num
		rightSum := sum - currentSum
		if leftSum == rightSum {
			return index
		}
	}

	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
}
