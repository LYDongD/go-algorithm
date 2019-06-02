package main

import "fmt"

func numTrees(n int) int {
    if n < 1 {
        return 0
    }

    count := make([]int, n+1)
    count[0], count[1] = 1, 1
    for i :=2 ; i <= n; i++ {
	for j := 1; j <= i; j++ {
	    count[i] += count[j-1] * count[i-j]
	}
    }

    return count[n]
}

func main() {
    fmt.Println(numTrees(3))
}

