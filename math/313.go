package main

import "fmt"

func nthSuperUglyNumber(n int, primes []int) int {
    if n <= 0 {
	panic("n should be bigger than 0")
    }

    if len(primes) == 0 {
	return 1
    }

    num := 2
    count := 1
    for  count < n {
	isUgly := tryAllFactors(num, transDic(primes))
	if isUgly {
	    count++
	    fmt.Println("ugly:", num, "count:", count)
	}
	num++
    }

    return num - 1
}

func tryAllFactors(num int, primesDic map[int]int) bool {
    if isPrime(num) && primesDic[num] == 0 {
	return false
    }

    hasTried := make(map[int]int)
    for i := 2; i < num; i++ {
	if hasTried[i] > 0 {
	    break
	}

	if num % i == 0 {
	    if isPrime(i) && primesDic[i] == 0 {
		return false
	    }

	    if isPrime(num / i) && primesDic[num / i] == 0 {
		return false
	    }

	    hasTried[num / i]++
	}
    }

    return true
}

func transDic(primes []int) map[int]int {
    result := make(map[int]int)
    for _, prime := range primes {
	result[prime]++
    }

    return result
}

func isPrime(num int) bool {
    for i := 2; i < num; i++ {
	if num % i == 0 {
	    return false
	}
    }

    return true
}

func main() {
    fmt.Println(nthSuperUglyNumber(12, []int{2,7,13,19}))
}
