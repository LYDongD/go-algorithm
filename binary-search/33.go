package main

import "fmt"

func search(nums []int, target int) int {
    if len(nums) == 0 {
	return -1
    }

    start := 0
    end := len(nums) - 1
    for start <= end {
	mid := start + (end - start) / 2
	if nums[mid] == target {
	    return mid
	}

	//确定mid的位置，左递增或右递增区间
	if nums[mid] >= nums[start] {
	    //确定target的位置
	    if target >= nums[start] && target <= nums[mid] {
		end = mid - 1
	    }else {
		start = mid + 1
	    }
	}else {
	    if target >= nums[mid] && target <= nums[end] {
		start = mid + 1
	    }else {
		end = mid - 1
	    }
	}
    }

    return -1
}

func main() {

    fmt.Println(search([]int{4,5,6,7,0,1,2}, 0))
    fmt.Println(search([]int{4,5,6,7,0,1,2}, 3))
}
