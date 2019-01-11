package array

import "fmt"

func findPairs(nums []int, k int) int {

	if len(nums) == 0 || k < 0 {
		return 0
	}

	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num] = countMap[num] + 1
	}

	fmt.Println(countMap)

	result := 0
	for num, count := range countMap {
		if k == 0 && count >= 2 {
			result++
		} else if k > 0 {
			numCouple := num + k
			_, ok := countMap[numCouple]
			if ok {
				result++
			}
		}
	}

	return result
}

func array() {

	fmt.Println(findPairs([]int{3, 1, 4, 1, 5}, 2))
	fmt.Println(findPairs([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(findPairs([]int{1, 3, 1, 5, 4}, 0))
	fmt.Println(findPairs([]int{0, 0, 0}, 0))
	fmt.Println(findPairs([]int{6, 2, 9, 3, 9, 6, 7, 7, 6, 4}, 3))
}
