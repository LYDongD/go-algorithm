package main

import "fmt"

func minCost(s string, cost []int) int {
	if len(s) <= 1 {
		return 0
	}

	minCost := 0
	start, end := -1, -1
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			if start == -1 {
				start = i
			}
		} else {
			if start >= 0 {
				end = i
				minCost += sumMinCost(cost, start, end)
				start = -1
				end = -1
			}
		}
	}

	if start >= 0 {
		end = len(s) - 1
		minCost += sumMinCost(cost, start, end)
	}

	return minCost
}

func sumMinCost(cost []int, start, end int) int {
	sum := 0
	max := cost[start]
	for i := start; i <= end; i++ {
		sum += cost[i]
		if max < cost[i] {
			max = cost[i]
		}
	}

	return sum - max
}

func main() {
	fmt.Println(minCost("abaac", []int{1, 2, 3, 4, 5}))
	fmt.Println(minCost("abc", []int{1, 2, 3}))
	fmt.Println(minCost("aabaa", []int{1, 2, 3, 4, 1}))
	fmt.Println(minCost("aaabbbabbbb", []int{3, 5, 10, 7, 5, 3, 5, 5, 4, 8, 1}))
	fmt.Println(minCost("aaaaaaaaaaaaaa", []int{1, 3, 6, 5, 4, 5, 4, 4, 2, 8, 3, 10, 6, 6}))
}
