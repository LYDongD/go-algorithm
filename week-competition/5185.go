package main

import "fmt"

func threeConsecutiveOdds(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	for i := 0; i <= len(arr)-3; i++ {
		if arr[i]%2 == 1 && arr[i+1]%2 == 1 && arr[i+2]%2 == 1 {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(threeConsecutiveOdds([]int{2, 6, 4, 1}))
	fmt.Println(threeConsecutiveOdds([]int{1, 2, 34, 3, 4, 5, 7, 23, 12}))
}
