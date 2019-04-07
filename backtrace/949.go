package main

import "fmt"
import "strconv"

func largestTimeFromDigits(A []int) string {
    if len(A) != 4 {
	return ""
    }

    maxTime := ""
    selected := ""
    rest := A
    return backtraceSelect(rest, selected, maxTime)
}

func backtraceSelect(rest []int, selected string, maxTime string) string {
    if len(rest) == 0 {
	if isTimeStr(selected) && selected > maxTime {
	    maxTime = selected
	}

	return maxTime
    }

    for index, num := range rest {
	copyRest := make([]int, len(rest))
	copy(copyRest, rest)
	copySelected := ""
	if len(selected) == 2 {
	    copySelected = selected + ":" + strconv.Itoa(num)
	}else {
	    copySelected = selected + strconv.Itoa(num)
	}
	maxTime = backtraceSelect(append(copyRest[:index], copyRest[index+1:]...), copySelected, maxTime)
    }

    return maxTime
}

func isTimeStr(str string) bool {
    if str[0] > '2' {
	return false
    }

    if str[0] == '2' && str[1] > '3' {
	return false
    }

    if str[3] > '5' {
	return false
    }

    return true
}

func main() {
    fmt.Println(largestTimeFromDigits([]int{1,2,3,4}))
    fmt.Println(largestTimeFromDigits([]int{5,5,5,5}))
}
