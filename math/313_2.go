package main

import "fmt"

func nthSuperUglyNumber(n int, primes []int) int {
    //defensive
    if n < 1 || len(primes) == 0 {
	return -1
    }

    //init uglys, products and idxes
    uglys := make([]int, n)
    uglys[0] = 1

    products, primeIdxes := make([]int, len(primes)), make([]int, len(primes))
    for i := 0; i < len(primes); i++ {
	products[i] = primes[i]
	primeIdxes[i] = 0
    }

    //find min product as next ugly
    next := 1
    for i := 0; i < n; i++ {
	uglys[i] = next
	next = int(^uint(0) >> 1)
	for j := 0; j < len(primes); j++ {
	    if products[j] == uglys[i] {
		primeIdxes[j] = primeIdxes[j] + 1
		products[j] = primes[j] * uglys[primeIdxes[j]]
	    }
	    if products[j] < next {
		next = products[j]
	    }
	}
    }

    fmt.Println(uglys)
    return uglys[n-1]
}

func main() {
    fmt.Println(nthSuperUglyNumber(12, []int{2,7,13,19}))
}
