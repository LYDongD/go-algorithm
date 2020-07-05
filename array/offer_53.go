package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	start, end := 0, len(nums)-1
	index := -1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			index = mid
			break
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	if index == -1 {
		return 0
	}

	count := 1
	for i := index - 1; i >= 0; i-- {
		if nums[i] == target {
			count++
		}
	}

	for j := index + 1; j <= len(nums)-1; j++ {
		if nums[j] == target {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 10))
	fmt.Println(search([]int{5, 7, 7, 8, 8, 10}, 6))
}
