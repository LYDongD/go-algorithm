package main

import "fmt"

func getWinner(arr []int, k int) int {
	if len(arr) == 0 {
		panic("arr is empty")
	}

	for t := 0; t <= len(arr)-1; t++ {
		if !allPrevLess(arr, t) {
			continue
		}

		count := 0
		if t > 0 {
			count = 1
		}

		if count == k {
			return arr[t]
		}

		hasBigger := false
		for i := t + 1; i <= len(arr)-1; i++ {
			if arr[i] < arr[t] {
				count++
				if count == k {
					return arr[t]
				}
			} else {
				hasBigger = true
				break
			}
		}

		if !hasBigger {
			return arr[t]
		}
	}

	panic("no winner")
}

func allPrevLess(arr []int, t int) bool {
	for i := 0; i < t; i++ {
		if arr[i] > arr[t] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(getWinner([]int{2, 1, 3, 5, 4, 6, 7}, 2))
	fmt.Println(getWinner([]int{3, 2, 1}, 10))
	fmt.Println(getWinner([]int{1, 9, 8, 2, 3, 7, 6, 4, 5}, 7))
	fmt.Println(getWinner([]int{1, 11, 22, 33, 44, 55, 66, 77, 88, 99}, 1000000000))
}
