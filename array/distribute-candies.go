package main

import "fmt"

func distributeCandies(candies []int) int {
	if len(candies) == 0 {
		return 0
	}

	//cal different num
	candyMap := make(map[int]int)
	for _, candy := range candies {
		candyMap[candy]++
	}

	candyMaxKinds := len(candyMap)
	candySelectCount := len(candies) / 2

	if candyMaxKinds < candySelectCount {
		return candyMaxKinds
	} else {
		return candySelectCount
	}
}

func main() {
	candies := []int{1, 1, 2, 2, 3, 3}
	candies2 := []int{1, 1, 2, 3}
	fmt.Println(distributeCandies(candies))
	fmt.Println(distributeCandies(candies2))
}
