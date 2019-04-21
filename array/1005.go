package main

import "fmt"

func largestSumAfterKNegations(A []int, K int) int {
    quickSort(A, 0, len(A)- 1)
    negativeCount := 0
    maxNegIndex := -1
    for i := 0; i < len(A); i++ {
	if A[i] >= 0 {
	    break
	}

	negativeCount++
	maxNegIndex = i
    }


    if negativeCount > K {
	for i := 0; i < K; i++ {
	    A[i] = -A[i]
	}
    }else {
	for i := 0; i < negativeCount; i++ {
	    A[i] = -A[i]
	}

	last := K - negativeCount
	if last % 2 == 1 {
	    if maxNegIndex == -1 {
		A[0] = -A[0]
	    }else if maxNegIndex == len(A) - 1 {
		A[maxNegIndex] = -A[maxNegIndex]
	    }else {
		if A[maxNegIndex] < A[maxNegIndex + 1] {
		    A[maxNegIndex] = -A[maxNegIndex]
		}else {
		    A[maxNegIndex+1] = -A[maxNegIndex+1] 
		}
	    }
	}
    }

    sum := 0
    for i := 0; i < len(A); i++ {
	sum += A[i]
    }

    return sum
}

func quickSort(A []int, start, end int) {
    if start >= end {
	return
    }

    pivotIndex := partion(A, start, end)
    quickSort(A, start, pivotIndex - 1)
    quickSort(A, pivotIndex + 1, end)
}

func partion(A []int, start, end int) int {
    pivot := A[end]
    pivotIndex := start
    for i := start; i < len(A) - 1; i++ {
	if A[i] < pivot {
	    A[i], A[pivotIndex] = A[pivotIndex], A[i]
	    pivotIndex++
	}
    }

    A[pivotIndex], A[end] = A[end], A[pivotIndex]
    return pivotIndex
}

func main() {
    //fmt.Println(largestSumAfterKNegations([]int{3,-1,0,2}, 3))
    fmt.Println(largestSumAfterKNegations([]int{-2,26,15,22,3,0,-6,11,-10,-4,12,14,-12,23,4,-7,30,3,-27,19,15,11,-29,-10,14,-28,3,24,-3,13,-28,30,24,13,16,21,1,-9,1,23,1,3,19,30,17,30,14,13,-5,22,16,-21,16,-29,7,0,-2,28,-21,24,21,-22,-7,-28,-6,25,10,0,-2,5,4,-9,21,-20,-26,24,15,-1,2,4,-9,-30,-9,18,-20,12,5,22,24,25,27,-30,1,24,-29,-14,5,21,-14,25}, 41))
}
