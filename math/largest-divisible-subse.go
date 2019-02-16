package main

import "fmt"
import "sort"

func largestDivisibleSubset(nums []int) []int {
    if len(nums) <= 1 {
	return nums
    }
    //sort
    sort.Ints(nums)

    //cal each num's common divisor count
    prev := make([]int, len(nums))
    for i := 0; i < len(prev); i++ {
	prev[i] = -1
    }

    counts := make([]int, len(nums))
    for i := 1; i < len(nums); i++ {
	maxCount := 0
	for j := 0; j < i; j++ {
	    if nums[i] % nums[j] == 0 && counts[j] + 1 > maxCount {
		maxCount = counts[j] + 1
		prev[i] = j
	    }
	}
	counts[i] = maxCount
    }

    //find num with max common divisor count
    bestNumIndex := 0
    maxCount := 0
    for index, count := range counts {
	if count > maxCount {
	    bestNumIndex = index
	    maxCount = count
	}
    }

    //get all common divisors by prev recursively
    var result []int
    return getPreDivisor(bestNumIndex, nums, prev, result)
}

func getPreDivisor(i int, nums []int, prev []int, result []int) []int {
    if prev[i] == -1 {
	result = append(result, nums[i])
	return result
    }

    result = getPreDivisor(prev[i], nums, prev, result)
    result = append(result, nums[i])
    return result
}

func main() {
    fmt.Println(largestDivisibleSubset([]int{1,2,4,8}))
    fmt.Println(largestDivisibleSubset([]int{1,2,3}))
    fmt.Println(largestDivisibleSubset([]int{546, 669}))
}
