package main

import "fmt"

func findJudge(N int, trust [][]int) int {
    if N == 1 {
	return 1
    }

    if N == 0 || len(trust) == 0 {
	return -1
    }


    trustedCountArr := make([]int, N+1)
    for _,trustItem := range trust {
	trustedCountArr[trustItem[0]]--
	trustedCountArr[trustItem[1]]++
    }

    for index, count := range trustedCountArr {
	if count == N - 1 {
	    return index
	}
    }

    return -1
}

func main() {
    fmt.Println(findJudge(2, [][]int{{1,2}}))
    fmt.Println(findJudge(3, [][]int{{1,3}, {2,3}}))
    fmt.Println(findJudge(4, [][]int{{1,3}, {2,3}, {1,4}, {2,4}, {4,3}}))
    
}
