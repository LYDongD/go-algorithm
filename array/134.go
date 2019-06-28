package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 || len(gas) != len(cost) {
		return -1
	}

	tank, current, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		current += gas[i] - cost[i]
		if current < 0 {
			start = i + 1
			current = 0
		}
	}

	if tank < 0 {
		return -1
	}

	return start

}

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
	fmt.Println(canCompleteCircuit([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}))
}
