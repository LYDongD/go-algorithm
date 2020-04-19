package main

import "fmt"

func twoSum(nums []int, target int) []int {
	result := []int{}
	if len(nums) <= 1 {
		return result
	}

	matchDic := make(map[int]int)
	for index, num := range nums {
		if val, ok := matchDic[num]; ok {
			result = append(result, val, index)
			return result
		}

		matchDic[target-num] = index
	}

	return result
}

func main() {
	fmt.Println(twoSum([]int{2}, 2))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
