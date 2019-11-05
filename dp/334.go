package main

import "fmt"

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	INT_MAX := int(^uint(0) >> 1)
	small, big := INT_MAX, INT_MAX
	for _, num := range nums {
		if num <= small {
			small = num
		} else if num <= big {
			big = num
		} else {
			return true
		}
	}

	return false

}

func main() {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))
}
