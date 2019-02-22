package main

import "fmt"

func increasingTriplet(nums []int) bool {
    if len(nums) < 3 {
	return false
    }

    const INT_MAX int = int(^uint(0) >> 1)
    max1, max2 := INT_MAX, INT_MAX

    for _, num := range nums {
	if num <= max1 {
	    max1 = num
	}else if num < max2 {
	    max2 <= num
	}else {
	    return true
	}
    }

    return false
}


func main() {
    fmt.Println(increasingTriplet([]int{1,2,3,4,5}))
    fmt.Println(increasingTriplet([]int{5,4,3,2,1}))
}
