package main

import "fmt"

func reversePairs(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	return mergeSort(nums, 0, len(nums)-1)

}

func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}

	mid := start + (end-start)/2
	count := mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end)
	left, right := start, mid+1
	tmp := []int{}
	//merge
	for left <= mid && right <= end {
		if nums[left] <= nums[right] {
			tmp = append(tmp, nums[left])
			count += right - (mid + 1)
			left++
		} else {
			tmp = append(tmp, nums[right])
			right++
		}
	}

	//handle left
	for i := left; i <= mid; i++ {
		tmp = append(tmp, nums[left])
		count += end - mid
	}

	//handle right
	for i := right; i <= end; i++ {
		tmp = append(tmp, nums[right])
	}

	//replace
	for i := start; i <= end; i++ {
		nums[i] = tmp[i-start]
	}

	return count
}

func main() {
	fmt.Println(reversePairs([]int{7, 5, 6, 4}))
}
