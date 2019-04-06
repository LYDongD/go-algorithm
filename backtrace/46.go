package main

import "fmt"

func permute(nums []int) [][]int {

    if len(nums) == 0 {
	return nil
    }

    result := [][]int{}
    selectedNums := []int{}
    restNums := nums
    return backtrace(selectedNums, restNums, result)
}

func backtrace(selectedNums []int, restNums []int, result [][]int) [][]int {
    if len(restNums) == 0 {
	result = append(result, selectedNums)
	return result
    }

    for index, num := range restNums {
	copySelectedNums := make([]int, len(selectedNums))
	copyRestNums := make([]int, len(restNums))
	copy(copySelectedNums, selectedNums)
	copy(copyRestNums, restNums)
	result = backtrace(append(copySelectedNums, num), append(copyRestNums[:index], copyRestNums[index+1:]...), result)
    }

    return result
}

func main() {
    fmt.Println(permute([]int{1,2,3}))
}
