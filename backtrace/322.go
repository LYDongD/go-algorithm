package main

import "fmt"
//import "sort"

func coinChange(coins []int, amount int) int {
    if len(coins) < 0 {
	return -1
    }

    if len(coins) == 0 && amount == 0 {
	return 0
    }

  //  sort.Sort(sort.Reverse(sort.IntSlice(coins)))
    minCount := int(^uint(0) >> 1)
    backtrace(coins, amount, 0, &minCount)
    if minCount == int(^uint(0) >> 1) {
	return -1
    }
    return minCount
}

func backtrace(coins []int, amount int, count int, minCount *int) {
    fmt.Println("amount", amount, "count", count)
    if amount == 0 {
	if *minCount > count {
	    *minCount = count
	}
    }

    if amount < 0 {
	return
    }

    for _, coin := range coins {
	backtrace(coins, amount - coin, count + 1, minCount)
    }
}


func main() {
    fmt.Println(coinChange([]int{1,2,5}, 11))
    fmt.Println(coinChange([]int{2}, 3))
    fmt.Println(coinChange([]int{186,419,83,408}, 6249))
}
