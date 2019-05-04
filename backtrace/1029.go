package main

import "fmt"

func twoCitySchedCost(costs [][]int) int {
    if len(costs) == 0 {
	return 0
    }

    minCost := int(^uint(0) >> 1)
    backtrace(costs, 0, 0, 0, 0, &minCost)

    return minCost
}

func backtrace(costs [][]int, k int, Anum, Bnum int, currentCost int, minCost *int) {
    if k == len(costs) {
	if currentCost < *minCost {
	    *minCost = currentCost
	}
    }

    //select A
    if Anum < len(costs) / 2 {
	backtrace(costs, k+1, Anum+1, Bnum, currentCost + costs[k][0], minCost)
    }

    //select B
    if Bnum < len(costs) / 2 {
	backtrace(costs, k+1, Anum, Bnum+1, currentCost + costs[k][1], minCost)
    }
}

func main() {
    fmt.Println(twoCitySchedCost([][]int{{10,20},{30,200},{400,50},{30,20}}))
}
