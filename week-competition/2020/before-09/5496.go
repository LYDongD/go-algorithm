package main

import "fmt"

func maxCoins(piles []int) int {
	if len(piles) == 0 || len(piles)%3 != 0 {
		return -1
	}

	quickSort(piles, 0, len(piles)-1)
	times := len(piles) / 3
	sum := 0
	for i := 0; i < times; i++ {
		sum += piles[i*2+1]
	}

	return sum
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	pivot := nums[end]
	split := start
	for i := start; i <= end-1; i++ {
		if nums[i] > pivot {
			swap(nums, i, split)
			split++
		}
	}

	swap(nums, split, end)

	quickSort(nums, start, split-1)
	quickSort(nums, split+1, end)
}

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

func main() {
	fmt.Println(maxCoins([]int{2, 4, 1, 2, 7, 8}))
	fmt.Println(maxCoins([]int{9, 8, 7, 6, 5, 1, 2, 3, 4}))
}
