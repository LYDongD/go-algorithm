package main

import "fmt"

func nthUglyNumber(n int) int {
    count := 1
    result := 1
    a,b,c := 2,3,5
    aIndex, bIndex, cIndex := 1,1,1
    for count < n {
	current := min(a * aIndex, b * bIndex, c * cIndex)
	if result < current {
	    result = current
	    count++
	}
	fmt.Println("count", count, "result:", result)
	if result == a * aIndex {
	    aIndex++
	}else if result == b * bIndex {
	    bIndex++
	}else {
	    cIndex++
	}
    }

    return result
}

func min(a,b,c int) int {
    min := a
    if b < min {
	min = b
    }

    if c < min {
	min = c
    }

    return min
}

func main() {
    fmt.Println(nthUglyNumber(11))
}
