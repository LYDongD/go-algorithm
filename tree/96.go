package main

import "fmt"

func numTrees(n int) int {
    if n < 1 {
	return 0
    }

    return countByDC(1,n)
}

func countByDC(start, end int) int {
    if start >= end {
	return 1
    }

    count := 0
    for i := start; i <= end; i++ {
	countLeft := countByDC(start, i-1)
	countRight := countByDC(i+1, end)
	count = count + countLeft * countRight
    }

    return count
}

func main() {
    fmt.Println(numTrees(1))
    fmt.Println(numTrees(3))
}
