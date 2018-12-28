package main

import "fmt"

func distributeCandies(candies []int) int {
	if len(candies) == 0 {
		return 0
	}

	//generate dp array
	n := len(candies)
	c := n / 2
	dp := make([][]int, n+1)
	for r := range dp {
		dp[r] = make([]int, c+1)
	}

	//init dp
	for r, row := range dp {
		for col := range row {
			if r == 0 {
				dp[r][col] = 0
			} else {
				dp[r][col] = -1
			}
		}
	}

	//store ele found
	foundMap := make(map[int]int)

	//fulfill dp
	for r := 1; r <= n; r++ {

		//do not select
		for col := range dp[r] {
			if dp[r-1][col] >= 0 {
				dp[r][col] = dp[r-1][col]
			}

		}

		//select
		for col := 0; col <= c-1; col++ {
			if dp[r-1][col] >= 0 {
				if foundMap[candies[r-1]] > 0 {
					if dp[r-1][col] > dp[r][col+1] {
						dp[r][col+1] = dp[r-1][col]
					}
				} else {
					if dp[r-1][col]+1 > dp[r][col+1] {
						dp[r][col+1] = dp[r-1][col] + 1
					}
				}
			}
		}

		//save the selected candy
		foundMap[candies[r-1]]++
	}

	return dp[n][c]
}

func main() {
	candies := []int{1, 1, 2, 2, 3, 3}
	candies2 := []int{1, 1, 2, 3}
	fmt.Println(distributeCandies(candies))
	fmt.Println(distributeCandies(candies2))
}
