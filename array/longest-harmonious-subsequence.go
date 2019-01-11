package array

import "fmt"

func findLHS(nums []int) int {

	n := len(nums)
	if n == 0 {
		return 0
	}

	mem := make(map[int]int)
	for _, num := range nums {
		mem[num]++
	}

	maxCount := 0
	count := 0
	for num := range mem {
		numCount, ok := mem[num+1]
		if ok {
			count = mem[num] + numCount
		}
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}

func array() {
	nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
	nums2 := []int{1, 2, 2, 1}
	fmt.Println(findLHS(nums2))
	fmt.Println(findLHS(nums))
}
