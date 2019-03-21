package main

import "fmt"

func fairCandySwap(A []int, B []int) []int {
    result := []int{}
    if len(A) == 0 || len(B) == 0 {
	return result
    }

    quickSort(A, 0, len(A)-1)
    quickSort(B, 0, len(B)-1)
    sumA, sumB := sum(A), sum(B)

    //two pointer
    cursorA, cursorB := 0, 0
    for cursorA < len(A) && cursorB < len(B) {
	tmpSumA := sumA - A[cursorA] + B[cursorB]
	tmpSumB := sumB - B[cursorB] + A[cursorA]
	if tmpSumA == tmpSumB {
	    return []int{A[cursorA], B[cursorB]}
	}else if tmpSumA > tmpSumB {
	    cursorA++
	}else {
	    cursorB++
	}
    }

    return result
}

func quickSort(A []int , start, end int){
    if start >= end {
	return
    }

    q := partion(A, start, end)
    quickSort(A, start, q - 1)
    quickSort(A, q + 1, end)
}

func partion(A []int, start, end int) int {
    pivot := A[end]
    q := start
    for i := start; i < end; i++ {
	if A[i] < pivot {
	    swap(A, i, q)
	    q++
	}
    }

    swap(A, q, end)
    return q
}

func swap(A []int, a, b int) {
    A[a], A[b] = A[b], A[a]
}

func sum(A []int) int{
    result := 0
    for _, num := range A {
	result += num
    }

    return result
}

func main() {
    //fmt.Println(fairCandySwap([]int{1,1}, []int{2,2}))
    //fmt.Println(fairCandySwap([]int{1,2}, []int{2,3}))
    //fmt.Println(fairCandySwap([]int{2}, []int{1,3}))
    //fmt.Println(fairCandySwap([]int{1,2,5}, []int{2,4}))
    A := []int{8,73,2,86,32}
    quickSort(A, 0, len(A) - 1)
    fmt.Println(A)
    fmt.Println(fairCandySwap([]int{8,73,2,86,32}, []int{56,5,67,100,31}))
}
