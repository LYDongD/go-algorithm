package main

import "fmt"

func minSwaps(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	n := len(grid)
	//cal continuos 0 count from right -> left
	zeroCounts := make([]int, n)
	for i := 0; i < n; i++ {
		count := 0
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 0 {
				count++
			} else {
				break
			}
		}
		zeroCounts[i] = count
	}

	//check 0 count by row, if not satisfied, swap it with another satisfied row
	swapCount := 0
	for i := 0; i < n; i++ {
		if zeroCounts[i] >= n-i-1 {
			continue
		}

		//greedy find and swap
		j := i + 1
		for ; j < n; j++ {
			if zeroCounts[j] >= n-i-1 {
				for k := j; k > i; k-- {
					zeroCounts[k], zeroCounts[k-1] = zeroCounts[k-1], zeroCounts[k]
					swapCount++
				}
				break
			}
		}

		//not found
		if j == n {
			return -1
		}
	}

	return swapCount
}

func main() {
	fmt.Println(minSwaps([][]int{{0, 0, 1}, {1, 1, 0}, {1, 0, 0}}))
	fmt.Println(minSwaps([][]int{{0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 0}}))
	fmt.Println(minSwaps([][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 1}}))
	fmt.Println(minSwaps([][]int{{1, 0, 0, 0, 0, 0}, {0, 0, 0, 1, 0, 0}, {0, 0, 0, 1, 0, 0}, {0, 1, 0, 0, 0, 0}, {0, 0, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 1}}))
}
