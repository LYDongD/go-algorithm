package main

import "fmt"

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return nums[0] == target
	}

	//binary search
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		fmt.Println(mid, nums[mid])
		if nums[mid] == target {
			return true
		}

		if nums[mid] > nums[end] {
			if target < nums[mid] && target >= nums[start] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else if nums[mid] < nums[end] {
			if target > nums[mid] && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			end--
		}

	}

	return false
}

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 5))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 6))
	fmt.Println(search([]int{3, 5, 1}, 3))
	fmt.Println(search([]int{1, 3, 5}, 1))
	fmt.Println(search([]int{1, 1, 3, 1}, 3))
}
