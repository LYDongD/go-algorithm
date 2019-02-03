package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start := 0
	end := len(nums)
	for start < end {
		middle := start + (end-start)/2
		if nums[middle] == target {
			return middle
		}

		if nums[middle] > target {
			end = middle
		} else {
			start = middle + 1
		}
	}

	return -1

}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 1}, 2))
}
