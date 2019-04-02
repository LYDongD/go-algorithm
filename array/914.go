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

    gcd := -1
    for _, count := range cardDic {
	if gcd == -1 {
	    gcd = count
	}else {
	    gcd = findGcd(gcd, count)
	}
    }

    return gcd >= 2
}

func findGcd(a, b int) int{
    divisor, divisored := -1, -1
    if a >= b {
	divisor = a
	divisored = b
    }else {
	divisor = b
	divisored = a
    }

    if divisored == 0 {
	return divisor
    }

    return findGcd(divisor % divisored, divisored)
}

func main() {
    fmt.Println(hasGroupsSizeX([]int{1,2,3,4,4,3,2,1}))
    fmt.Println(hasGroupsSizeX([]int{1,1,1,2,2,2,3,3}))
    fmt.Println(hasGroupsSizeX([]int{1,1,1,1,2,2,2,2,2,2}))
}
