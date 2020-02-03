package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := []int{}
	if len(nums1) == 0 {
		nums = nums2
	} else if len(nums2) == 0 {
		nums = nums1
	} else {
		i, j := 0, 0
		for i < len(nums1) || j < len(nums2) {
			if i >= len(nums1) {
				nums = append(nums, nums2[j])
				j++
			} else if j >= len(nums2) {
				nums = append(nums, nums1[i])
				i++
			} else if nums1[i] < nums2[j] {
				nums = append(nums, nums1[i])
				i++
			} else {
				nums = append(nums, nums2[j])
				j++
			}
		}
	}

	mid := float64(nums[len(nums)/2])
	if len(nums)%2 == 0 {
		mid = (mid + float64(nums[len(nums)/2-1])) / 2.0
	}

	return mid
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
