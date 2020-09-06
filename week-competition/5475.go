package main

import "fmt"

func countGoodTriplets(arr []int, a int, b int, c int) int {
	if len(arr) < 3 {
		return 0
	}

	count := 0
	for i := 0; i <= len(arr)-3; i++ {
		for j := i + 1; j <= len(arr)-2; j++ {
			for k := j + 1; k <= len(arr)-1; k++ {
				if isGoodTurple(arr, i, j, k, a, b, c) {
					count++
				}
			}
		}
	}

	return count
}

func isGoodTurple(arr []int, i, j, k, a, b, c int) bool {
	return abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}

	return a
}

func main() {
	fmt.Println(countGoodTriplets([]int{3, 0, 1, 1, 9, 7}, 7, 2, 3))
	fmt.Println(countGoodTriplets([]int{1, 1, 2, 2, 3}, 0, 0, 1))
}
