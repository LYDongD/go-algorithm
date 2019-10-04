package main

import "fmt"

func maxProfit(prices []int) int {
    if len(prices) == 0 {
	return 0
    }

    n := len(prices)
    rest, buy, sell := make([]int, n), make([]int, n), make([]int, n)
    rest[0] = 0
    buy[0] = -prices[0]
    sell[0] = ^(int(^uint(0) >> 1))
    for i := 1; i < len(prices); i++ {
	rest[i] = max(rest[i-1], sell[i-1])
	buy[i] = max(buy[i-1], rest[i-1] - prices[i])
	sell[i] = buy[i-1] + prices[i]
    }

    return max(rest[n-1], sell[n-1])
}

func max(a, b int) int {
    if a > b {
	return a
    }

    return b
}

func main(){
    fmt.Println(maxProfit([]int{1,2,3,0,2}))
}
