package main

import "fmt"

func validMountainArray(A []int) bool {
    if len(A) < 3 {
	return false
    }

    peakCount := 0
    for i := 1; i < len(A) - 1; i++ {
	if A[i] > A[i-1] && A[i] > A[i+1] {
	    peakCount++
	    if peakCount > 1 {
		return false
	    }
	}else if A[i] == A[i-1] || A[i] == A[i+1] {
	    return false
	}
    }

    //edge count
    last := len(A) - 1
    if (A[0] > A[1] || A[last] > A[last-1]) && peakCount >= 1 {
	return false
    }

    return peakCount == 1
}

func main() {
    fmt.Println(validMountainArray([]int{2,1}))
    fmt.Println(validMountainArray([]int{3,5,5}))
    fmt.Println(validMountainArray([]int{0,3,2,1}))
    fmt.Println(validMountainArray([]int{0,1,2,1,2}))
    fmt.Println(validMountainArray([]int{2,1,2,3,5,7,9,10,12,14,15,16,18,14,13}))
    fmt.Println(validMountainArray([]int{14,82,89,84,79,70,70,68,67,66,63,60,58,54,44,43,32,28,26,25,22,15,13,12,10,8,7,5,4,3}))
}
