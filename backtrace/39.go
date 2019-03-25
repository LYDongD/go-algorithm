package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	if len(candidates) == 0 {
		return result
	}

	quickSort(candidates, 0, len(candidates)-1)
	return backtrace(candidates, target, 0, []int{}, result)
}

func backtrace(candidates []int, target, sum int, combination []int, result [][]int) [][]int {
	if sum == target {
		result = append(result, combination)
		return result
	}

	for _, candidate := range candidates {
		if len(combination) > 0 && combination[len(combination)-1] > candidate {
			continue
		}

		if sum+candidate <= target {
			tmpCombination := addByConmbination(combination, candidate)
			result = backtrace(candidates, target, sum+candidate, tmpCombination, result)
		}
	}

	return result
}

func addByConmbination(combination []int, candidate int) []int {
	result := make([]int, len(combination)+1)
	for i, num := range combination {
		result[i] = num
	}

	result[len(result)-1] = candidate
	return result

}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}

	q := partion(nums, start, end)
	quickSort(nums, start, q-1)
	quickSort(nums, q+1, end)
}

func partion(nums []int, start, end int) int {
	pivort := nums[end]
	q := start
	for i := start; i < end; i++ {
		if nums[i] < pivort {
			swap(nums, i, q)
			q++
		}
	}

	swap(nums, q, end)
	return q
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{7, 3, 2}, 18))
}
