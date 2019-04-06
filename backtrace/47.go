package main

import "fmt"

func permuteUnique(nums []int) [][]int {
    result := [][]int{}

    if len(nums) == 0 {
	return result
    }

    if len(nums) == 1 {
	return append(result, nums)
    }

    selected := []int{}
    rest := nums
    return backtrace(selected, rest, result)
}

func backtrace(selected []int, rest []int, result [][]int) [][]int {
    if len(rest) == 0 {
	result = append(result, selected)
	return result
    }

    existedMap := make(map[int]int)
    for index, num := range rest {
	if existedMap[num] > 0 {
	    continue
	}else {
	    existedMap[num]++
	}

	copySelected, copyRest := make([]int, len(selected)), make([]int, len(rest))
	copy(copySelected, selected)
	copy(copyRest, rest)
	result = backtrace(append(copySelected, num), append(copyRest[:index], copyRest[index+1:]...), result)
    }

    return result
}

func main() {
    fmt.Println(permuteUnique([]int{1,1,2}))
}
