package main

import "fmt"

func isMonotonic(A []int) bool {
    if len(A) == 0 {
	return false
    }

    if len(A) == 1 {
	return true
    }

    flag := 0
    for i := 1; i < len(A); i++ {
	if A[i] > A[i-1] {
	    if flag == -1 {
		return false
	    }

	    flag = 1
	}else if A[i] < A[i-1] {
	    if flag == 1 {
		return false
	    }

	    flag = -1
	}
    }

    return true
}


func main() {
    fmt.Println(isMonotonic([]int{1,2,2,3}))
    fmt.Println(isMonotonic([]int{6,5,4,4}))
    fmt.Println(isMonotonic([]int{1,3,2}))
    fmt.Println(isMonotonic([]int{1,2,4,5}))
    fmt.Println(isMonotonic([]int{1,1,1,1}))
}
