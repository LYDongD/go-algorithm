package main

import "fmt"

func numSquares(n int) int {
    if n <=0 {
	return 0
    }

    return countSquare(n)
}

func countSquare(n int) int {
    if isSquare(n) {
	return 1
    }

    i := 1
    num := 0
    min := int(^uint(0) >> 1)
    for {
	t := i * i
	if t > n {
	    break
	}

	left := countSquare(n - t)
	if left > 0 {
	    num = 1 + left
	}else {
	    continue
	}

	if num < min {
	    min = num
	}
	i++
    }

    if min == int(^uint(0) >> 1) {
	min = 0
    }
    return min
}

func isSquare(n int) bool {
    if n == 1 {
	return true
    }

    start, end := 1, n/2
    for start <= end {
	mid := start + (end - start) / 2
	if n / mid == mid && n % mid == 0 {
	    return true
	}else if n / mid > mid || (n / mid == mid && n % mid > 0) {
	    start = mid + 1
	}else {
	    end = mid - 1
	}
    }

    return false
}

func main() {
    fmt.Println(numSquares(62))
    fmt.Println(numSquares(12))
}
