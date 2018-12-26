package main

import "fmt"

func matrixReshape(nums [][]int, r int, c int) [][]int {

	m := len(nums)
	n := len(nums[0])

	if m*n != r*c {
		return nums
	}

	s := make(chan int, r*c)

	go walkMatrix(nums, s)

	result := make([][]int, r)
	for i := range result {
		result[i] = make([]int, c)
	}

	for i, row := range result {
		for j := range row {
			result[i][j] = <-s
		}
	}

	return result

}

func walkMatrix(nums [][]int, s chan int) {
	for _, row := range nums {
		for _, cell := range row {
			s <- cell
		}
	}
}

func main() {

	nums := [][]int{{1, 2}, {3, 4}}
	nums2 := [][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(matrixReshape(nums, 1, 4))
	fmt.Println(matrixReshape(nums2, 2, 3))
}
