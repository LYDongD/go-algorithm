package main

import "fmt"

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	recA := abs(C-A) * abs(D-B)
	recB := abs(G-E) * abs(H-F)

	left := max(A, E)
	right := min(C, G)
	bottom := max(B, F)
	top := min(D, H)

	intersectionArea := 0
	if left < right && bottom < top {
		intersectionArea = (right - left) * (top - bottom)
	}

	return recA + recB - intersectionArea
}

func abs(a int) int {
	if a >= 0 {
		return a
	}

	return -a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2))
}
