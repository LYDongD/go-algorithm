package main

import "fmt"

func numTriplets(nums1 []int, nums2 []int) int {
	if len(nums1) < 2 || len(nums2) < 2 {
		return 0
	}

	count := 0
	for _, num1 := range nums1 {
		square := num1 * num1
		for index2, num2 := range nums2 {
			if square%num2 == 0 {
				target := square / num2
				for i := index2 + 1; i < len(nums2); i++ {
					if nums2[i] == target {
						count++
					}
				}
			}
		}
	}

	for _, num2 := range nums2 {
		square := num2 * num2
		for index1, num1 := range nums1 {
			if square%num1 == 0 {
				target := square / num1
				for i := index1 + 1; i < len(nums1); i++ {
					if nums1[i] == target {
						count++
					}
				}
			}
		}
	}
	return count
}

func main() {
	fmt.Println(numTriplets([]int{7, 4}, []int{5, 2, 8, 9}))
	fmt.Println(numTriplets([]int{1, 1}, []int{1, 1, 1}))
	fmt.Println(numTriplets([]int{7, 7, 8, 3}, []int{1, 2, 9, 7}))
	fmt.Println(numTriplets([]int{4, 7, 9, 11, 23}, []int{3, 5, 1024, 12, 18}))
}
