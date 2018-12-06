package main

import "fmt"

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return nil
	}

	result := make([]int, 2)
	targetMap := make(map[int]int)
	for i, num := range nums {
		k, ok := targetMap[num]
		if !ok {
			targetMap[target-num] = i
		} else {
			result[0] = k
			result[1] = i
		}
	}

	return result

}

func main() {

	nums := []int{2, 7, 11, 5}
	target := 9
	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSum(nums, 16))
}
