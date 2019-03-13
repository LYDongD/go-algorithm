package main

import "fmt"

func peakIndexInMountainArray(A []int) int {
    if len(A) == 0 {
	return -1
    }

    left := 0
    right := len(A) - 1
    for left < right {
	mid := left + (right - left) / 2
	if A[mid] < A[mid+1] {
	    left = mid + 1
	}else {
	    right = mid
	}
    }

    return right

}

func main() {
    fmt.Println(peakIndexInMountainArray([]int{0,1,0}))
    fmt.Println(peakIndexInMountainArray([]int{0,2,1,0}))
}
