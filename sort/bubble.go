package main

import "fmt"

func bubbleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	n := len(nums)
	for i := 0; i < n; i++ {
		hasSwap := false
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				swap(nums, j, j+1)
				hasSwap = true
			}
		}

		if !hasSwap {
			break
		}
	}
}

func swap(nums []int, a, b int) {
	tmp := nums[a] ^ nums[b]
	nums[a] = tmp ^ nums[a]
	nums[b] = tmp ^ nums[b]
}

func main() {
	//nums := []int{6, 5, 7, 3, 1, 3}
	nums2 := []int{1, 2, 3, 4, 5, 6, 7}
	bubbleSort(nums2)
	fmt.Println(nums2)
}
