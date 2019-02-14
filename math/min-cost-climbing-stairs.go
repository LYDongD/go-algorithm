package main

import "fmt"

func minCostClimbingStairs(cost []int) int {

    n := len(cost)
    if n == 0 {
	return 0
    }

    minCostCache := make(map[int]int)
    cost1 := minCost(n-1, cost, minCostCache)
    cost2 := minCost(n-2, cost, minCostCache)
    if cost1 < cost2 {
	return cost1
    }

    return cost2
}

func minCost(n int, cost []int, cache map[int]int) int {
    if n <= 1 {
	return cost[n]
    }

    var cost1 int
    if cache[n-1] > 0 {
	cost1 = cost[n] + cache[n-1]
    }else {
	cost1 = cost[n] + minCost(n-1, cost, cache)
    }

    var cost2 int
    if cache[n-2] > 0 {
	cost2 = cost[n] + cache[n-2] 
    }else {
	cost2 = cost[n] + minCost(n-2, cost, cache)
    }

    if cost1 < cost2 {
	cache[n] = cost1
	return cost1
    }

    cache[n] = cost2
    return cost2
} 

func main() {

    fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
    fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))


}
