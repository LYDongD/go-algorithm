package main

import "fmt"

func reachNumber(target int) int {
	if target == 0 {
		return 0
	}

	k := 1
	cache := make(map[int]int)
	for {
		move := reachNumberRecursively(target, k, 0, cache)
		if move < int(^uint(0)>>1) {
			return move
		}
		k++
	}

	return -1
}

func reachNumberRecursively(target int, k int, move int, cache map[int]int) int {
	if k == 0 {
		if target == 0 {
			return move
		}

		return int(^uint(0) >> 1)
	}

	if cache[target] > 0 {
		return cache[target]
	}

	move1 := reachNumberRecursively(target+k, k-1, move+1, cache)
	move2 := reachNumberRecursively(target-k, k-1, move+1, cache)
	minMove := min(move1, move2)
	if minMove < int(^uint(0)>>1) {
		cache[target] = minMove
	}
	return minMove
}

func min(a int, b int) int {
	if a <= b {
		return a
	}

	return b
}

func main() {
	fmt.Println(reachNumber(3))
	fmt.Println(reachNumber(2))
}
