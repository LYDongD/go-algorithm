package main

import "fmt"

func insertSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	n := len(nums)
	for i := 1; i < n; i++ {
		j := i - 1
		num := nums[i]
		for ; j >= 0; j-- {
			if nums[j] > num {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}

		nums[j+1] = num
	}
}

func main() {
	nums := []int{7, 3, 5, 8, 3, 1}
	insertSort(nums)
	fmt.Println(nums)
}
