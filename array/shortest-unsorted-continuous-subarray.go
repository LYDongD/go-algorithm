package array

import "fmt"

func findUnsortedSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sortedNums := sort(nums)
	start := -1
	end := -1
	for i := 0; i < len(nums); i++ {
		if sortedNums[i] != nums[i] {
			if start < 0 {
				start = i
			} else {
				end = i
			}
		}
	}

	if start < 0 {
		return 0
	} else {
		return end - start + 1
	}

}

func sort(nums []int) []int {
	sortedNums := make([]int, len(nums))
	for i, num := range nums {
		sortedNums[i] = num
	}

	//quick sort
	quickSort(sortedNums, 0, len(sortedNums)-1)
	return sortedNums
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
	splitIndex := start
	for i := start; i <= end-1; i++ {
		if nums[i] < pivot {
			swap(nums, splitIndex, i)
			splitIndex++
		}
	}

	swap(nums, splitIndex, end)
	return splitIndex
}

func swap(nums []int, a int, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

func array() {
	nums := []int{2, 6, 4, 8, 10, 9, 15}
	nums2 := []int{1, 2, 3, 4}
	nums3 := []int{2, 1}
	nums5 := []int{1, 2, 5, 3, 4}
	fmt.Println(findUnsortedSubarray(nums))
	fmt.Println(findUnsortedSubarray(nums2))
	fmt.Println(findUnsortedSubarray(nums3))
	fmt.Println(findUnsortedSubarray(nums5))
}
