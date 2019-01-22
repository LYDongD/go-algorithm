package main

import "fmt"

func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	max := 0
	current := 1
	for i := 0; i < n-1; i++ {
		if nums[i+1] > nums[i] {
			current++
		} else {
			if current > max {
				max = current
			}
			current = 1
		}
	}

	if current > max {
		max = current
	}

	return max
}

func main() {
	nums1 := []int{1, 3, 5, 4, 7}
	fmt.Println(findLengthOfLCIS(nums1))
	nums2 := []int{2, 2, 2, 2, 2}
	fmt.Println(findLengthOfLCIS(nums2))
}
