package main

import "fmt"
import "sort"

func numMovesStones(a int, b int, c int) []int {

	stones := []int{a, b, c}
	sort.Ints(stones)

	if stones[2]-stones[0] == 2 {
		return []int{0, 0}
	}

	minDelta := min(stones[1]-stones[0], stones[2]-stones[1])
	minMove := 2
	maxMove := stones[2] - stones[0] - 2
	if minDelta == 1 || minDelta == 2 {
		minMove = 1
	}

	return []int{minMove, maxMove}

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	fmt.Println(numMovesStones(1, 2, 5))
	fmt.Println(numMovesStones(4, 3, 2))
	fmt.Println(numMovesStones(3, 5, 1))
	fmt.Println(numMovesStones(1, 7, 2))
}
