package array

import "fmt"

func nextGreaterElement(findNums []int, nums []int) []int {
	if len(findNums) == 0 || len(nums) == 0 {
		return nil
	}

	//cache nums index
	numMap := make(map[int]int)

	for index, num := range nums {
		numMap[num] = index
	}

	result := make([]int, len(findNums))
	//find index from cache, then find the first greater num
	for index, findNum := range findNums {
		j, ok := numMap[findNum]
		if ok {
			isFound := false
			for k := j + 1; k < len(nums); k++ {
				if nums[k] > findNum {
					result[index] = nums[k]
					isFound = true
					break
				}
			}

			if !isFound {
				result[index] = -1
			}
		}
	}

	return result
}

func array() {

	findNums := []int{2, 4}
	nums := []int{1, 2, 3, 4}

	fmt.Println(nextGreaterElement(findNums, nums))
}
