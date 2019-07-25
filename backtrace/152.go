package main

import "fmt"

var max int

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		panic("nums is empty")
	}

	max = ^(int(^uint(0) >> 1))
	for i, num := range nums {
		if num > max {
			max = num
		}
		backtrace(nums, i+1, num)
	}
	return max
}

func backtrace(nums []int, selectIndex, current int) {
	if selectIndex == len(nums) {
		return
	}

	if current*nums[selectIndex] > max {
		max = current * nums[selectIndex]
	}

	backtrace(nums, selectIndex+1, current*nums[selectIndex])
}

func main() {
	fmt.Println(maxProduct([]int{0, 2}))
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
}
