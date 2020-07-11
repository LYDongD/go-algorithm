package main

import "fmt"

//find first ele equal 2 target
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)>>1
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}

			end = mid - 1
		}
	}

	return -1
}

//find last ele less than target
func search2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)>>1
		if nums[mid] > target {
			mid = end - 1
		} else {
			//check if it's the last one less than target
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid
			}

			start = mid + 1
		}
	}

	return -1
}

//find ele in a circular ordered array
func search3(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	head, tail := nums[0], nums[len(nums)-1]
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)>>1
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			//may be on smaller part
			if nums[mid] > tail && target <= tail {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			//may be on bigger part
			if nums[mid] < head && target >= head {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}

	return -1
}

func main() {
	fmt.Println(search3([]int{4, 5}, 4))
	fmt.Println(search3([]int{4, 5, 4}, 4))
	fmt.Println(search3([]int{4, 5}, 5))
	fmt.Println(search3([]int{4, 5, 6, 1, 2, 3, 4}, 5))
	fmt.Println(search3([]int{4, 5, 6, 1, 2, 3, 4}, 2))
	fmt.Println(search3([]int{4, 5, 6, 1, 2, 3, 4}, 2))
	//fmt.Println(search([]int{2, 3, 3, 3, 3, 3, 5, 7}, 3))
	//fmt.Println(search2([]int{2, 3, 3, 3, 3, 3, 5, 7}, 3))
	//fmt.Println(search([]int{2, 3, 3, 3, 3, 3, 5, 7}, 6))
	//fmt.Println(search2([]int{2, 3, 3, 3, 3, 3, 5, 7}, 6))
}
