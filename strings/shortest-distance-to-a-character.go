package main

import "fmt"

func shortestToChar(S string, C byte) []int {
    var result []int
    if len(S) == 0 {
	return result
    }

    if len(S) == 1 && S[0] == C {
	return []int{0}
    }

    //cal C indexs
    CIndexs := []int{}
    for index, char := range S {
	if byte(char) == C {
	    CIndexs = append(CIndexs, index)
	}
    }

    //cal min diff
    for index := range S {
	result = append(result, minDis(index, CIndexs))
    }

    return result
}

func minDis(index int, CIndexs []int) int {
    min := int(^uint(0) >> 1)
    for _, cIndex := range CIndexs {
	if abs(index - cIndex) < min {
	    min = abs(index - cIndex)
	}
    }

    return min
}

func abs(a int) int {
    if a >= 0 {
	return a
    }

    return -1 * a
}

func main() {
    fmt.Println(shortestToChar("loveleetcode", 'e'))
}
