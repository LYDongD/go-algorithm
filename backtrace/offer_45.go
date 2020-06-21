package main

import "fmt"
import "strconv"

func minNumber(nums []int) string {
    result := ""
    if len(nums) == 0 {
	return result
    }

    quickSort(nums, 0, len(nums) - 1)
    for _, num := range nums {
	result = result + strconv.Itoa(num)
    }

    return result

}

func quickSort(nums []int, start, end int) {
    if start >= end {
	return
    }

    pivot := nums[end]
    split := start
    for i := start; i <= end; i++ {
	if small2Pivot(nums[i], pivot) {
	    swap(nums, i, split)
	    split++
	}
    }

    //swap pivot
    swap(nums, end, split)
    quickSort(nums, start, split - 1)
    quickSort(nums, split+1, end)
}

func small2Pivot(num, pivot int) bool {
    numStr, pivotStr := strconv.Itoa(num), strconv.Itoa(pivot)
    return numStr + pivotStr < pivotStr + numStr
}

func swap(nums []int, a, b int) {
    nums[a], nums[b] = nums[b], nums[a]
}

func main(){
    fmt.Println(minNumber([]int{3,30,34,5,9}))
    fmt.Println(minNumber([]int{1,1,1}))
}
