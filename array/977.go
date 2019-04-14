package main

import "fmt"

func sortedSquares(A []int) []int {
	if len(A) == 0 {
		return nil
	}

	result := make([]int, len(A))
	start, end := 0, len(A)-1
	for i := range A {
		startSquare, endSquare := A[start]*A[start], A[end]*A[end]
		if startSquare > endSquare {
			result[len(A)-i-1] = startSquare
			start++
		} else {
			result[len(A)-i-1] = endSquare
			end--
		}
	}

	return result
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}
