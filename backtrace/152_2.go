package main

import "fmt"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		panic("nums is empty")
	}

	if len(nums) == 1 {
		return nums[0]
	}

	max, min := []int{nums[0]}, []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		max = append(max, findMax(max[i-1]*num, min[i-1]*num, num))
		min = append(min, findMin(max[i-1]*num, min[i-1]*num, num))
	}

	return findMaxEle(max)
}

func findMax(a, b, c int) int {
	max := a
	if max < b {
		max = b
	}

	if max < c {
		max = c
	}

	return max
}

func findMin(a, b, c int) int {
	min := a
	if min > b {
		min = b
	}

	if min > c {
		min = c
	}

	return min
}

func findMaxEle(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

func main() {
	fmt.Println(maxProduct([]int{-1, 2, -9, -6}))
	fmt.Println(maxProduct([]int{0, 2}))
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
}
