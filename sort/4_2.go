package main

import "fmt"

//binary search partionX to find median
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//make first array shorter one
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	//binary search at nums1
	m, n := len(nums1), len(nums2)
	start, end := 0, m
	const MAX_VAL = int(^uint(0) >> 1)
	const MIN_VAL = ^MAX_VAL
	for start <= end {
		//mid as partion index but not array index
		partionX := start + (end-start)/2
		partionY := (m+n+1)/2 - partionX

		//partion
		maxLeftX := compareAndReturn(partionX, 0, MIN_VAL, nums1, partionX-1)
		minRightX := compareAndReturn(partionX, m, MAX_VAL, nums1, partionX)
		maxLeftY := compareAndReturn(partionY, 0, MIN_VAL, nums2, partionY-1)
		minRightY := compareAndReturn(partionY, n, MAX_VAL, nums2, partionY)

		//found
		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (m+n)%2 == 0 {
				return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2.0
			} else {
				return float64(max(maxLeftX, maxLeftY))
			}
		} else if maxLeftX > minRightY {
			end = partionX - 1
		} else {
			start = partionX + 1
		}
	}

	panic("arguments illegal")
}

func compareAndReturn(source, target, val1 int, nums []int, index int) int {
	if source == target {
		return val1
	}

	return nums[index]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
