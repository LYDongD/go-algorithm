package main

import "fmt"

func subsets(nums []int) [][]int {
	result := [][]int{{}}
	if len(nums) == 0 {
		return result
	}

	result = backtrace(0, nums, []int{}, result)

	return result
}

func backtrace(start int, nums []int, selecedArr []int, result [][]int) [][]int {
	if start == len(nums) {
		return result
	}

	//select 1 ele from start
	for i := start; i < len(nums); i++ {
		tmpSelectedArr := make([]int, len(selecedArr))
		copy(tmpSelectedArr, selecedArr)
		tmpSelectedArr = append(tmpSelectedArr, nums[i])
		result = append(result, tmpSelectedArr)
		result = backtrace(i+1, nums, tmpSelectedArr, result)
	}

	return result
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
