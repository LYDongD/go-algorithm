package main

import "fmt"

func binarySearch(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	start := 0
	end := len(arr) - 1
	for start < end {
		middle := start + (end-start)/2
		if arr[middle] < target {
			start = middle + 1
		} else {
			end = middle
		}
	}

	if arr[start] == target {
		return start
	}

	return -1
}

func main() {
	fmt.Println(binarySearch([]int{1, 3, 11, 11, 11, 23, 3}, 11))
	fmt.Println(binarySearch([]int{1, 3, 11, 11, 11, 23, 3}, 6))
}
