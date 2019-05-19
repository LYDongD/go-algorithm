package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	consecutiveCount := 1
	last := nums[0]
	lastIndex := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != last {
			consecutiveCount = 1
			last = nums[i]
		} else {
			last = nums[i]
			consecutiveCount++
			if consecutiveCount > 2 {
				continue
			}
		}

		//swap goal ele to left
		nums[lastIndex], nums[i] = nums[i], nums[lastIndex]
		lastIndex++
	}
	return lastIndex
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3}))
	fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 2, 3, 3}))
}
