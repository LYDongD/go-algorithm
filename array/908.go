package main

import "fmt"

func smallestRangeI(A []int, K int) int {
    if len(A) == 0 {
	return 0
    }

    min, max := A[0], A[0]
    for _, num := range A {
	if num > max {
	    max = num
	}else if num < min {
	    min = num
	}
    }

    if max - min - 2*K >= 0 {
	return max - min - 2*K
    }

    return 0
}

func main() {
    fmt.Println(smallestRangeI([]int{1}, 0))
    fmt.Println(smallestRangeI([]int{0,10}, 2))
    fmt.Println(smallestRangeI([]int{1,3,6}, 3))
}
