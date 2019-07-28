package main

import "fmt"

func findMin(nums []int) int {
    if len(nums) == 0 {
	panic("nums is empty")
    }

    if len(nums) == 1 {
	return nums[0]
    }


    start, end := 0, len(nums) - 1
    if nums[end] > nums[0] {
	return nums[0]
    }

    for start < end {
	mid := start + (end - start ) / 2
	if nums[mid] > nums[mid+1] {
	    return nums[mid+1]
	}

	if nums[mid] < nums[mid - 1] {
	    return nums[mid]
	}

	if nums[mid] > nums[0] {
	    start = mid + 1
	}else {
	    end  = mid - 1
	}
    }

    return nums[0]

}

func main() {
    fmt.Println(findMin([]int{3,1,2}))
    fmt.Println(findMin([]int{1,2,3}))
    fmt.Println(findMin([]int{3,4,5,1,2}))
    fmt.Println(findMin([]int{4,5,6,7,0,1,2}))
}
