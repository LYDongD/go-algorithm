package main

import "fmt"

func hammingDistance(x int, y int) int {
    if x == y {
	return 0
    }

    //transfer differ to 1
    z := x ^ y

    //count 1 bit
    count := 0
    for z > 0 {
	if z & 1 == 1 {
	    count++
	}

	z = z >> 1
    }

    return count
}

func main() {
    fmt.Println(hammingDistance(1,4))
}
