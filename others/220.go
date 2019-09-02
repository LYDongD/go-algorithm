package main

import "fmt"

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) <= 1 || k <= 0 || t < 0 {
		return false
	}

	minVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if minVal > nums[i] {
			minVal = nums[i]
		}
	}

	bucketSize := t + 1
	bucketDic := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		bucketIndex := (nums[i] - minVal) / bucketSize
		if _, ok := bucketDic[bucketIndex]; ok {
			return true
		}

		bucketDic[bucketIndex] = nums[i]
		if left, ok := bucketDic[bucketIndex-1]; ok {
			if abs(nums[i]-left) <= t {
				return true
			}
		}

		if right, ok := bucketDic[bucketIndex+1]; ok {
			if abs(nums[i]-right) <= t {
				return true
			}
		}

		if i >= k {
			delete(bucketDic, (nums[i-k]-minVal)/bucketSize)
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
	fmt.Println(containsNearbyAlmostDuplicate([]int{2, 0, -2, 2}, 2, 1))
	fmt.Println(containsNearbyAlmostDuplicate([]int{-3, 3}, 2, 4))
	fmt.Println(containsNearbyAlmostDuplicate([]int{2, 2}, 3, 0))
	fmt.Println(containsNearbyAlmostDuplicate([]int{-1, -1}, 1, 0))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2, 3, 1}, 3, 0))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 0, 1, 1}, 1, 2))
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 5, 9, 1, 5, 9}, 2, 3))
}
