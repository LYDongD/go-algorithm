package main

import "fmt"
import "sort"

func subsetsWithDup(nums []int) [][]int {
    result := [][]int{{}}
    if len(nums) == 0 {
	return result
    }

    sort.Ints(nums)

    return backtrace(nums, result, 0, []int{})
}

func backtrace(nums []int, result [][]int, start int, selected []int) [][]int{
    if start == len(nums) {
	return result
    }

    //select from start, avoid to select previous
    for i := start; i < len(nums); i++ {
	//avoid duplicated select
	if i > start && nums[i] == nums[i-1] {
	    continue
	}

	tmpSelected := make([]int, len(selected))
	copy(tmpSelected, selected)
	tmpSelected = append(tmpSelected, nums[i])
	result = append(result, tmpSelected)
	result = backtrace(nums, result, i + 1, tmpSelected)
    }

    return result
}

func main() {
    fmt.Println(subsetsWithDup([]int{4,4,4,1,4}))
}
