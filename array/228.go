package main

import "fmt"
import "strconv"

func summaryRanges(nums []int) []string {
    result := []string{}
    if len(nums) == 0 {
	return result
    }

    start, end := nums[0], nums[0]
    for i := 0; i < len(nums); i++ {
	if i == len(nums) - 1 || nums[i+1] - nums[i] > 1 {
	    end = nums[i]
	    if start == end {
		result = append(result, strconv.Itoa(start))
	    }else {
		result = append(result, strconv.Itoa(start) + "->" + strconv.Itoa(end))
	    }

	    if i < len(nums) - 1 {
		start = nums[i+1]
	    }
	}
    }

    return result
}

func main() {
    fmt.Println(summaryRanges([]int{0,1,2,4,5,7}))
    fmt.Println(summaryRanges([]int{0,2,3,4,6,8,9}))
}
