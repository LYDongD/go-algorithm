ackage main

import "fmt"

func combinationSum3(k int, n int) [][]int {
    result := [][]int{}
    if k <= 0 || n <= 0 {
	return result
    }

    result = backtraceSelect(1, k, n, 0, []int{}, [][]int{})

    return result
}

func backtraceSelect(start, k, n, sum int, selected []int, result [][]int) [][]int{
    if len(selected) == k {
	if sum  == n {
	    result = append(result, selected)
	}
	return result
    }

    for i := start; i <= 9; i++ {
	if sum + i > n {
	    return result
	}

	tmp := make([]int, len(selected))
	copy(tmp, selected)
	tmp = append(tmp, i)
	result = backtraceSelect(i+1, k, n, sum + i, tmp, result)
    }

    return result
}

func main() {
    fmt.Println(combinationSum3(3,7))
    fmt.Println(combinationSum3(3,9))
}
