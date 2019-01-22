package main

import "fmt"

func threeSum(nums []int) [][]int {

	//sort
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	//3sum -> 2 sum
	fixedNumMap := make(map[int]int)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		num := nums[i]

		//avoid duplication
		_, ok := fixedNumMap[num]
		if ok {
			continue
		} else {
			fixedNumMap[num]++
		}

		twoSum := 0 - num
		fmt.Println(num, twoSum)
		twoSumMap := make(map[int]int)

		//find two sum and get a full triple
		duplicateMap := make(map[int]int)
		for j := i + 1; j < len(nums); j++ {
			current := nums[j]
			//find two sum
			next, ok := twoSumMap[current]
			fmt.Println("查表", current, next, ok)
			if ok {
				triple := make([]int, 3)
				triple[0] = num
				triple[1] = next
				triple[2] = current
				result = append(result, triple)
				//avoid duplicate
				delete(twoSumMap, current)
			} else {
				fmt.Println("建表", twoSum-current, current)
				//avoid duplicate
				if duplicateMap[twoSum-current] == 0 {
					twoSumMap[twoSum-current] = current
				}
				duplicateMap[twoSum-current]++
			}
		}
	}

	return result
}

func quickSort(nums []int, start int, end int) {

	//divide and conqure
	if start < end {
		q := partion(nums, start, end)
		quickSort(nums, start, q-1)
		quickSort(nums, q+1, end)
	}
}

func partion(nums []int, start int, end int) int {
	pivort := nums[end]
	insertIndex := start
	for i := start; i < end; i++ {
		if nums[i] < pivort {
			swap(nums, insertIndex, i)
			insertIndex++
		}
	}
	swap(nums, insertIndex, end)
	return insertIndex
}

func swap(nums []int, a int, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

func main() {
	nums := []int{0, 2, 2, 3, 0, 1, 2, 3, -1, -4, 2}
	fmt.Println(threeSum(nums))
}
