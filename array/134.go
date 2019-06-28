package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 || len(gas) != len(cost) {
		return -1
	}

	//candidate from 0 -> len(cost) - 1
	result := -1
	for start := 0; start < len(cost); start++ {

		leftGas := 0
		hasFound := true
		steps := 0
		next := start
		for steps < len(cost) {
			leftGas = leftGas + gas[next] - cost[next]
			if leftGas < 0 {
				hasFound = false
				break
			}
			steps++
			next++
			if next > len(cost)-1 {
				next = 0
			}
		}

		if hasFound {
			result = start
			break
		}
	}

	return result

}

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
	fmt.Println(canCompleteCircuit([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}))
}
