package main

import "fmt"

func searchRange(nums []int, target int) []int {
    result := []int{-1, -1}
    if len(nums) == 0 {
	return result
    }

    //find min index target
    start, end := 0, len(nums) - 1
    for start < end {
	mid := start + (end - start) / 2 
	if nums[mid] < target {
	    start = mid + 1
	}else {
	    end = mid
	}
    }

    if nums[start] == target {
	result[0] = start
    }

    //fine max index target
    start, end = 0, len(nums) - 1
    for start < end {
	mid := start + (end - start) / 2 
	if nums[mid] <= target {
	    start = mid + 1
	}else {
	    end = mid
	}
    }

    if nums[end] == target {
	result[1] = end
    }else if start > 0 && nums[start - 1] == target {
	result[1] = end -1
    }

    return result
}

func main() {
    fmt.Println(searchRange([]int{1}, 0))
    fmt.Println(searchRange([]int{1}, 1))
    fmt.Println(searchRange([]int{2,2}, 1))
    fmt.Println(searchRange([]int{2,2}, 2))
    fmt.Println(searchRange([]int{5,7,7,8,8,10}, 8))
    fmt.Println(searchRange([]int{5,7,7,8,8,10}, 6))
}

