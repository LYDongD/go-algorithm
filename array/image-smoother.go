package main

import "fmt"

func imageSmoother(M [][]int) [][]int {
	m := len(M)

	if m == 0 {
		return nil
	}

	n := len(M[0])

	result := make([][]int, m)
	for row := range result {
		result[row] = make([]int, n)
	}

	for i, rowArr := range result {
		for j := range rowArr {
			sum := 0
			count := 0
			for rn := i - 1; rn <= i+1; rn++ {
				for cn := j - 1; cn <= j+1; cn++ {
					if rn >= 0 && rn < m && cn >= 0 && cn < n {
						sum += M[rn][cn]
						count++
					}
				}
			}
			result[i][j] = sum / count
		}
	}

	return result
}

func main() {
	nums := [][]int{{2, 3, 4}, {5, 6, 7}, {8, 9, 10}, {11, 12, 13}, {14, 15, 16}}
	fmt.Println(imageSmoother(nums))
}
