package main

import "fmt"

func binaryGap(N int) int {
    if N == 0 {
	return 0
    }

    index := 0
    lastOne, currentOne := 0, 0
    maxDis := 0
    for N > 0 {
	bit := N % 2
	index++
	if bit == 1 {
	    currentOne = index
	    if lastOne > 0 {
		maxDis = max(currentOne - lastOne, maxDis)
	    }

	    lastOne = index
	}

	N = N / 2
    }

    return maxDis
}

func max(a, b int) int {
    if a > b {
	return a
    }

    return b
}

func main() {
    fmt.Println(binaryGap(5))
    fmt.Println(binaryGap(22))
    fmt.Println(binaryGap(6))
    fmt.Println(binaryGap(8))
}
