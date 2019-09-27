package main

import "fmt"

func findDuplicate(nums []int) int {
	if len(nums) <= 1 {
		return -1
	}

	//binary search
	start, end := 1, len(nums)-1
	for start < end {
		mid := start + (end-start)/2
		count := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				count++
			}
		}

		if count > mid {
			end = mid
		} else {
			start = mid + 1
		}

	}

	return start
}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
}
