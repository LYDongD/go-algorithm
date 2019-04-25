package main

import "fmt"

func numPairsDivisibleBy60(time []int) int {
    if len(time) <= 1 {
	return 0
    }

    //分析法
    modDic := make([]int, 60)
    count := 0
    for _, timeItem := range time {
	modA := timeItem % 60
	modB := (60 - modA) % 60
	count += modDic[modB]
	modDic[modA]++
    }

    return count
}

func main() {
    fmt.Println(numPairsDivisibleBy60([]int{30,20,150,100,40}))
    fmt.Println(numPairsDivisibleBy60([]int{60,60,60}))
}
