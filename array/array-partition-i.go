package main

import "fmt"

func arrayPairSum(nums []int) int {

	//quick sort
	sort(nums)

	sum := 0
	for index, num := range nums {
		if index%2 == 0 {
			sum += num
		}
	}

	return sum
}

func sort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
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
	storeIndex := start
	for i := start; i <= end-1; i++ {
		if nums[i] < pivot {
			swap(nums, i, storeIndex)
			storeIndex++
		}
	}

	swap(nums, storeIndex, end)
	return storeIndex
}

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	nums := []int{1, 4, 3, 2, 7, 9}
	fmt.Println(arrayPairSum(nums))
}
