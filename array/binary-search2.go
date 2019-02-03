package main

import "fmt"

func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start := 0
	end := len(nums) - 1
	for start <= end {
		middle := start + (end-start)/2

		//find the max index
		if nums[middle] == target {
			for i := middle + 1; i <= end; i++ {
				if nums[i] > target {
					return i - 1
				}
			}
		}

		if nums[middle] > target {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}

	//find the max ele less then target
	for i := start; i >= 0; i-- {
		if nums[i] < target {
			return i
		}
	}

	return -1
}

func main() {

	fmt.Println(binarySearch([]int{1, 3, 5, 7, 9}, 7))
	fmt.Println(binarySearch([]int{1, 3, 5, 7, 9}, 8))
	fmt.Println(binarySearch([]int{1, 3, 3, 3, 5, 6}, 3))
}
