package main

import "fmt"

func diStringMatch(S string) []int {
    if len(S) == 0 {
	return nil
    }

    N := len(S)
    increaseNum, decreaseNum := 0, N

    result := make([]int, N+1)
    for index, char := range S {
	current := 0
	if char == 'I' {
	    current = increaseNum
	    increaseNum++
	}else if char == 'D' {
	    current = decreaseNum
	    decreaseNum--
	}
	result[index] = current
    }

    result[N] = increaseNum

    return result
}

func main() {
    fmt.Println(diStringMatch("IDID"))
    fmt.Println(diStringMatch("DDI"))
    fmt.Println(diStringMatch("III"))
}
