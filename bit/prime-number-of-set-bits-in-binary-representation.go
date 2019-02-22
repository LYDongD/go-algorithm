package main

import "fmt"

func countPrimeSetBits(L int, R int) int {
	if L > R {
		return -1
	}

	count := 0
	for i := L; i <= R; i++ {
		if isPrime(oneBitCount(i)) {
			count++
		}
	}

	return count
}

func oneBitCount(n int) int {
	count := 0
	for n > 0 {
		if n&1 == 1 {
			count++
		}

		n = n >> 1
	}

	return count
}

func isPrime(n int) bool {
	if n <= 3 {
		return n-1 != 0
	}

	for i := 2; i <= sqrt(n); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func sqrt(n int) int {
	start := 1
	end := n
	for start < end {
		middle := start + (end-start)/2

		square := middle * middle
		if square == n {
			return middle
		}

		if square > n {
			end = middle
		} else {
			start = middle + 1
		}
	}

	return start - 1
}

func main() {
	fmt.Println(countPrimeSetBits(6, 10))
	fmt.Println(countPrimeSetBits(10, 15))
}
