package main

import "fmt"

func twoCitySchedCost(costs [][]int) int {
	quickSortByDelta(costs, 0, len(costs)-1)

	sum := 0
	N := len(costs) / 2
	for i := 0; i < N; i++ {
		sum = sum + costs[i][1] + costs[2*N-i-1][0]
	}

	return sum
}

func quickSortByDelta(costs [][]int, start, end int) {
	if start >= end {
		return
	}

	pivotIndex := partion(costs, start, end)
	quickSortByDelta(costs, start, pivotIndex-1)
	quickSortByDelta(costs, pivotIndex+1, end)
}

func partion(costs [][]int, start, end int) int {
	pivotDelta := costs[end][1] - costs[end][0]
	pivotIndex := start
	for i := start; i < end; i++ {
		currentDelta := costs[i][1] - costs[i][0]
		if currentDelta < pivotDelta {
			swap(costs, pivotIndex, i)
			pivotIndex++
		}
	}

	swap(costs, pivotIndex, end)
	return pivotIndex
}

func swap(costs [][]int, a, b int) {
	costs[a], costs[b] = costs[b], costs[a]
}

func main() {
	fmt.Println(twoCitySchedCost([][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}))
}
