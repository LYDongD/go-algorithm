package main

import "fmt"

func largeGroupPositions(S string) [][]int {
    result := [][]int{}
    if len(S) == 0 {
	return result
    }

    count := 1
    for i := 1; i < len(S); i++ {
	if S[i] == S[i-1] {
	    count++
	}else {
	    if count >= 3 {
		result = append(result, []int{i - count, i - 1})
	    }
	    count = 1
	}
    }

    if count >= 3 {
	n := len(S)
	result = append(result, []int{n-count, n-1})
    }

    return result
}

func main() {
    fmt.Println(largeGroupPositions("abbxxxxzzy"))
    fmt.Println(largeGroupPositions("abc"))
    fmt.Println(largeGroupPositions("abcdddeeeeaabbbcd"))
}
