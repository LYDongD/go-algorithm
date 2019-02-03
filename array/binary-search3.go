package main

import "fmt"

func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start := 0
	end := len(nums) - 1
	for start < end {
		middle := start + (end-start)/2
		if nums[middle] > target {
			end = middle
		} else {
			start = middle + 1
		}
	}

	return start - 1
}

func main() {

	fmt.Println(binarySearch([]int{1, 3, 5, 7, 9}, 7))
	fmt.Println(binarySearch([]int{1, 3, 5, 7, 9}, 8))
	fmt.Println(binarySearch([]int{1, 3, 3, 3, 5, 6}, 3))
}
