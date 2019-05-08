package main

import "fmt"

func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}

	last0, first2 := 0, len(nums)-1
	for i := 0; i <= first2; i++ {
		if nums[i] == 0 {
			swap(nums, i, last0)
			last0++
		} else if nums[i] == 2 {
			swap(nums, i, first2)
			first2--
			i--
		}

	}
}

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	nums2 := []int{1, 2, 0}
	nums3 := []int{2, 1, 2}
	sortColors(nums)
	sortColors(nums2)
	sortColors(nums3)
	fmt.Println(nums)
	fmt.Println(nums2)
	fmt.Println(nums3)
}
