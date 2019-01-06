package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {

	if n == 0 {
		return true
	}

	if handleEdgePos(flowerbed, &n) {
		return true
	}

	count := countCanPlacePos(flowerbed)

	if n <= count {
		return true
	}

	return false
}

func handleEdgePos(flowerbed []int, n *int) bool {

	if len(flowerbed) == 1 && flowerbed[0] == 0 {
		return *n <= 1
	}

	if flowerbed[0] == 0 && flowerbed[1] == 0 {
		flowerbed[0] = 1
		*n = *n - 1
		if *n == 0 {
			return true
		}
	}

	if flowerbed[len(flowerbed)-1] == 0 && flowerbed[len(flowerbed)-2] == 0 {
		flowerbed[len(flowerbed)-1] = 1
		*n = *n - 1
		if *n == 0 {
			return true
		}
	}

	return false
}

func countCanPlacePos(flowerbed []int) int {
	count := 0
	consecutiveZeros := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			consecutiveZeros++
		} else {
			if consecutiveZeros > 2 && consecutiveZeros%2 == 0 {
				count += consecutiveZeros/2 - 1
			} else if consecutiveZeros > 2 && consecutiveZeros%2 == 1 {
				count += consecutiveZeros / 2
			}

			consecutiveZeros = 0
		}
	}

	return count

}

func main() {
	flowers := []int{1, 0, 0, 0, 0}
	fmt.Println(canPlaceFlowers(flowers, 2))
}
