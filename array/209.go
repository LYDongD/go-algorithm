package main

import "fmt"

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	minLength := len(nums) + 1
	for i := 0; i < len(nums); i++ {
		if nums[i] >= s {
			return 1
		}
		sum := nums[i]
		length := 1
		for k := i + 1; k < len(nums); k++ {
			sum = sum + nums[k]
			length++
			if sum >= s && length < minLength {
				minLength = length
				break
			}
		}
	}

	if minLength == len(nums)+1 {
		minLength = 0
	}
	return minLength
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(15, []int{1, 2, 3, 4, 5}))
}
