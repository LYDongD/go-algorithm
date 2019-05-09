package main

import "fmt"

func combine(n int, k int) [][]int {
	if n <= 0 || k > n {
		return nil
	}

	result := [][]int{}
	selected := []int{}
	result = backtrace(n, k, 0, selected, result)

	return result
}

func backtrace(n, k, s int, selected []int, result [][]int) [][]int {

	//selected all
	if k == 0 {
		result = append(result, selected)
		return result
	}

	//select from s to n
	for i := s + 1; i <= n; i++ {
		currentSelected := make([]int, len(selected))
		copy(currentSelected, selected)
		currentSelected = append(currentSelected, i)
		result = backtrace(n, k-1, i, currentSelected, result)
	}

	return result
}

func main() {
	fmt.Println(combine(4, 2))
	fmt.Println(combine(5, 3))
}
