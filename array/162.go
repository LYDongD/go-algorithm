package main

import "fmt"

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	start, end := 0, len(nums)-1

	for start < end {
		mid := start + (end-start)/2

		if nums[start] > nums[start+1] {
			return start
		}

		if nums[end] > nums[end-1] {
			return end
		}

		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}

		if nums[mid] < nums[mid+1] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
	fmt.Println(findPeakElement([]int{3, 4, 3, 2, 1}))
}
