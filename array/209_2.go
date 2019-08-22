package main

import "fmt"

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	minLength := len(nums) + 1
	left := 0
	currentSum := 0
	//move right
	for right := 0; right < len(nums); right++ {
		currentSum += nums[right]
		for currentSum >= s {
			minLength = min(minLength, right-left+1)
			//move left
			currentSum -= nums[left]
			left++

			if minLength == 1 {
				return 1
			}
		}
	}

	if minLength == len(nums)+1 {
		minLength = 0
	}

	return minLength
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	fmt.Println(minSubArrayLen(15, []int{1, 2, 3, 4, 5}))
}
