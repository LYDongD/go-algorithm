package main

import "fmt"

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	//tails: save LTS by vaired length
	tails := []int{}
	for _, num := range nums {
		size := len(tails)
		if size == 0 || num > tails[size-1] {
			tails = append(tails, num)
			continue
		}

		//find the first ele bigger than num and switch it
		start, end := 0, size-1
		for start < end {
			mid := start + (end-start)/2
			if tails[mid] < num {
				start = mid + 1
			} else {
				end = mid
			}
		}

		tails[start] = num
	}
	return len(tails)
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 4}))
}
