package main

import "fmt"

func findK(nums []int, k int) int {
	if len(nums) < k {
		panic("k must not excess nums' bound")
	}

	return findKRecursively(k, nums, 0, len(nums)-1)
}

func findKRecursively(k int, nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}

	pivot := nums[end]
	j := start
	for i := start; i < end; i++ {
		if nums[i] < pivot {
			swap(nums, i, j)
			j++
		}
	}

	swap(nums, j, end)

	//find next range
	if j+1 == k {
		return nums[j]
	} else if j < k {
		return findKRecursively(k, nums, j+1, end)
	} else {
		return findKRecursively(k, nums, start, j-1)
	}

}

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

func main() {
	fmt.Println(findK([]int{6, 1, 3, 4, 7, 2, 2, 5}, 3))
	fmt.Println(findK([]int{6, 1, 3, 4, 7, 5}, 3))
}
