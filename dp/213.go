package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	bestAmount := 0
	best := make([]int, len(nums))
	//dp from 0 -> n - 1
	best[0] = nums[0]
	best[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums)-1; i++ {
		best[i] = max(best[i-1], best[i-2]+nums[i])
	}
	bestAmount = best[len(nums)-2]

	//dp from 1 -> n
	best[1] = nums[1]
	best[2] = max(nums[1], nums[2])
	for i := 3; i < len(nums); i++ {
		best[i] = max(best[i-1], best[i-2]+nums[i])
	}
	bestAmount = max(bestAmount, best[len(nums)-1])

	return bestAmount
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
}
