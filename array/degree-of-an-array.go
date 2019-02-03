package main

import "fmt"

type Frequency struct {
	count int
	start int
	end   int
}

func findShortestSubArray(nums []int) int {
	//defensive
	if len(nums) <= 1 {
		return len(nums)
	}

	//generate frequency map
	freMap := make(map[int]*Frequency)
	for index, num := range nums {
		frequency, ok := freMap[num]
		if !ok {
			freMap[num] = &Frequency{count: 1, start: index, end: index}
		} else {
			frequency.count = frequency.count + 1
			frequency.end = index
		}
	}

	//find maxCount and minSub
	maxCount := 0
	minSub := -1
	for _, fre := range freMap {
		if fre.count > maxCount {
			maxCount = fre.count
			minSub = fre.end - fre.start + 1
		} else if fre.count == maxCount {
			subLen := fre.end - fre.start + 1
			if subLen < minSub {
				minSub = subLen
			}
		}
	}

	return minSub
}

func main() {
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1, 4, 2}))
}
