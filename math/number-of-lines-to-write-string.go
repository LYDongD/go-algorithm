package main

import "fmt"

func numberOfLines(widths []int, S string) []int {
    if len(S) == 0 {
	return []int{0,0}
    }

    lineWidth := 0
    lineCount := 0
    for _, char := range S {
	if lineWidth + widths[char - 'a'] <= 100 {
	    lineWidth = lineWidth + widths[char - 'a']
	}else {
	    lineWidth = widths[char - 'a']
	    lineCount++
	}
    }

    return []int{lineCount + 1, lineWidth}
}

func main() {
    fmt.Println(numberOfLines([]int{10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10},"abcdefghijklmnopqrstuvwxyz") )
    fmt.Println(numberOfLines([]int{4,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10}, "bbbcccdddaaa"))
}
