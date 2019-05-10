package main

import "fmt"
import "math"

func powerfulIntegers(x int, y int, bound int) []int {

	if x == 1 && y == 1 {
		if bound > 1 {
			return []int{2}
		}
	}

	result := []int{}
	dic := make(map[int]int)
	i := 0
	for pow(x, i) < bound {
		j := 0
		for {
			sum := pow(x, i) + pow(y, j)
			if sum > bound {
				break
			}

			if dic[sum] == 0 {
				result = append(result, sum)
				dic[sum]++
			}

			if y == 1 {
				break
			}

			j++
		}

		if x == 1 {
			break
		}

		i++
	}

	return result
}

func pow(x, y int) int {
	floatX, floatY := float64(x), float64(y)
	return int(math.Pow(floatX, floatY))
}

func main() {
	fmt.Println(powerfulIntegers(1, 1, 1))
	fmt.Println(powerfulIntegers(2, 3, 2))
	fmt.Println(powerfulIntegers(2, 1, 10))
	fmt.Println(powerfulIntegers(2, 3, 10))
	fmt.Println(powerfulIntegers(3, 5, 15))
}
