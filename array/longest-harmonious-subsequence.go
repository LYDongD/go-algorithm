package main

import "fmt"

func findLHS(nums []int) int {

	n := len(nums)
	if n == 0 {
		return 0
	}

	quickSort(nums, 0, n-1)

	maxCount := 0
	for i := 0; i < n; i++ {
		harmonuousCount := 0
		equalCount := 1
		for j := i + 1; j < n; j++ {
			if nums[j]-nums[i] == 0 {
				equalCount++
			} else if nums[j]-nums[i] == 1 {
				if equalCount > 0 {
					harmonuousCount = equalCount
					equalCount = 0
				}
				harmonuousCount++
			} else { //diff > 1 should break
				break
			}

		}

		if harmonuousCount > maxCount {
			maxCount = harmonuousCount
		}
	}

	return maxCount
}

func quickSort(nums []int, start int, end int) {
	if start >= end {
		return
	}

	q := partion(nums, start, end)
	quickSort(nums, start, q-1)
	quickSort(nums, q+1, end)
}

func partion(nums []int, start int, end int) int {
	pivot := nums[end]
	pivotIndex := start
	for i := start; i < end; i++ {
		if nums[i] < pivot {
			//swap the unknown ele to the right, the less one to the left
			swap(nums, i, pivotIndex)
			pivotIndex++
		}
	}

	swap(nums, pivotIndex, end)
	return pivotIndex
}

func swap(nums []int, a int, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

func main() {
	//	nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	nums2 := []int{1, 2, 2, 1}
	fmt.Println(findLHS(nums2))
}
