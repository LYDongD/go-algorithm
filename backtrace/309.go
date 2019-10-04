package main

import "fmt"

func maxProfit(prices []int) int {
    if len(prices) == 0 {
	return 0
    }

    max := 0
    backtrace(prices, 0, 0, 2, &max)
    return max
}

func backtrace(prices []int, day int, profit int, lastOperation int, max *int) {
    if day == len(prices)  {
	if profit > *max {
	    *max = profit
	}
	return
    }

    price := prices[day]

    if lastOperation == 1 {//last sell, then cool down
	backtrace(prices, day+1, profit, 2, max)
    }else if lastOperation == 0 { //last buy, then sell, next cooldown
	//sell or wait
	backtrace(prices, day+1, profit + price, 1, max)
	backtrace(prices, day+1, profit, 0, max)
    }else {
	//buy or wait
	backtrace(prices, day+1, profit - price, 0, max)
	backtrace(prices, day+1, profit, 2, max)
    }
}

func main() {
    fmt.Println(maxProfit([]int{1,2,4}))
    fmt.Println(maxProfit([]int{1,2,3,0,2}))
}
