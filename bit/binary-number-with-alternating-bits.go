package main

import "fmt"

func hasAlternatingBits(n int) bool {

	last := n % 2
	n = n / 2
	for n > 0 {
		if last == n%2 {
			return false
		}

		last = n % 2
		n = n / 2
	}

	return true
}

func main() {
	fmt.Println(hasAlternatingBits(5))
	fmt.Println(hasAlternatingBits(11))
	fmt.Println(hasAlternatingBits(10))
}
