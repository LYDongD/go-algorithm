package main

import "fmt"

func sortArrayByParity(A []int) []int {
    if len(A) <= 1 {
	return A
    }

    start, end := 0, len(A) - 1
    for start < end {
	if A[start] % 2 == 0 && A[end] % 2 == 0 {
	    start++
	}else if A[start] % 2 == 0 && A[end] % 2 == 1 {
	    start++
	    end--
	}else if A[start] % 2 == 1 && A[end] % 2 == 0 {
	    A[start], A[end] = A[end], A[start]
	    start++
	    end--
	}else {
	    end--
	}
    }

    return A
}

func main() {
    fmt.Println(sortArrayByParity([]int{3,1,2,4}))
}
