package main

import "fmt"

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {

	if k < 0 || t < 0 || len(nums) == 0 {
		return false
	}

	for start := 0; start < len(nums); start++ {
		for j := 1; j <= k; j++ {
			if start+j >= len(nums) {
				break
			}

			diff := abs(nums[start+j] - nums[start])
			if diff <= t {
				return true
			}
		}
	}

	return false
}

func abs(a int) int {
	if a >= 0 {
		return a
	}

	return -a
}

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 0, 1, 1}, 1, 2))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
}
