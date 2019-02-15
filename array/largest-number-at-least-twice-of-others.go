package main

import "fmt"

func dominantIndex(nums []int) int {
    if len(nums) < 1 {
	return -1
    }

    if len(nums) == 1 {
        return 0
    }

    max, maxIndex := nums[0], 0
    secondMax, secondMaxIndex := -1, -1

    for i := 1; i < len(nums); i++ {
	if nums[i] > max {
	    secondMax, secondMaxIndex = max, maxIndex
	    max, maxIndex = nums[i], i
	}else if nums[i] > secondMax {
	    secondMax, secondMaxIndex = nums[i], i
	}
    }

    _ = secondMaxIndex

    if max >= 2 * secondMax {
	return maxIndex
    }

    return -1
}

func main() {
    fmt.Println(dominantIndex([]int{3, 6, 1, 0}))
    fmt.Println(dominantIndex([]int{1,2,3,4}))
}

