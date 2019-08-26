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

	return max(robRange(nums, 0, len(nums)-2), robRange(nums, 1, len(nums)-1))
}

func robRange(nums []int, i, j int) int {
	best := make([]int, len(nums))
	best[i] = nums[i]
	best[i+1] = max(best[i], nums[i+1])
	for k := i + 2; k <= j; k++ {
		best[k] = max(best[k-1], best[k-2]+nums[k])
	}
	return best[j]
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
