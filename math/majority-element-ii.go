package main

import "fmt"

func majorityElement(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	num1, num2, count1, count2 := 0, 0, 0, 0

	for _, num := range nums {

		if num1 == num {
			count1++
		} else if num2 == num {
			count2++
		} else if count1 == 0 {
			num1 = num
			count1++
		} else if count2 == 0 {
			num2 = num
			count2++
		} else {
			count1--
			count2--
		}
	}

	count1, count2 = 0, 0
	for _, num := range nums {
		if num == num1 {
			count1++
		} else if num == num2 {
			count2++
		}
	}

	var result []int
	if count1 > len(nums)/3 {
		result = append(result, num1)
	}

	if count2 > len(nums)/3 {
		result = append(result, num2)
	}

	return result
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{1, 1, 1, 3, 3, 2, 2, 2}))
}
