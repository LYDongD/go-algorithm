package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
    result := [][]int{}
    if len(candidates) == 0 {
	return result
    }

    return backtrace(candidates, target, 0, []int{}, result)
}

func backtrace(candidates []int, target, sum int, combination []int, result [][]int) [][]int {
    if sum == target {
	result = append(result, combination)
	return result
    }

    if sum > target {
	return result
    }

    for _, candidate := range candidates {
	if combination[len(combination) - 1] > candidate {
	    continue
	}

	result = backtrace(candidates, target, sum + candidate, append(combination, candidate), result)
    }

    return result
}

func main() {
    fmt.Println(combinationSum([]int{2,3,6,7}, 7))
    fmt.Println(combinationSum([]int{2,3,5}, 8))
}
