package main

import "fmt"

func hasGroupsSizeX(deck []int) bool {
    if len(deck) <= 1 {
	return false
    }

    cardDic := make(map[int]int)
    for _, num := range deck {
	cardDic[num]++
    }

    groupCount := len(deck)
    for _,count := range cardDic {
	if count < groupCount {
	    groupCount = count
	}
    }

    if groupCount <= 1 {
	return false
    }

    for _, count := range cardDic {
	if count % groupCount != 0 {
	    return false
	}
    }

    return true
}

func main() {
    fmt.Println(hasGroupsSizeX([]int{1,2,3,4,4,3,2,1}))
    fmt.Println(hasGroupsSizeX([]int{1,1,1,2,2,2,3,3}))
    fmt.Println(hasGroupsSizeX([]int{1,1,2,2,2,2}))
}
