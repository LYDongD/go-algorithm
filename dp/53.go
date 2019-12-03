package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = max(sums[i-1]+nums[i], nums[i])
	}

	maxSum := sums[0]
	for i := 1; i < len(sums); i++ {
		if sums[i] > maxSum {
			maxSum = sums[i]
		}
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{0}))
}
