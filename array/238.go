package main

import "fmt"

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	if len(nums) <= 1 {
		return result
	}

	//init left product
	leftProducts := make([]int, len(nums))
	leftProducts[0] = 1
	for i := 0; i < len(nums)-1; i++ {
		leftProducts[i+1] = nums[i] * leftProducts[i]
	}

	//cal result
	rightProduct := 1
	for i := len(leftProducts) - 1; i >= 0; i-- {
		result[i] = leftProducts[i] * rightProduct
		rightProduct = rightProduct * nums[i]
	}

	return result
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
