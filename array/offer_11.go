package main

import "fmt"

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		panic("numbers is empty")
	}

	start, end := 0, len(numbers)-1
	for start < end {
		mid := start + (end-start)/2
		//right
		if numbers[mid] < numbers[end] {
			end = mid
		} else if numbers[mid] > numbers[end] {
			start = mid + 1
		} else {
			end--
		}
	}

	return numbers[start]
}

func main() {
	fmt.Println(minArray([]int{3, 4, 5, 1, 2}))
	fmt.Println(minArray([]int{2, 2, 2, 0, 1}))
	fmt.Println(minArray([]int{3, 1, 1}))
	fmt.Println(minArray([]int{0, 1, 2}))
	fmt.Println(minArray([]int{3, 3, 1, 3}))
}
