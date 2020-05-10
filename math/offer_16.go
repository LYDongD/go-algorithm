package main

import "fmt"

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	nBit := n
	if nBit < 0 {
		nBit = -nBit
	}

	result := float64(1.0)
	for nBit > 0 {
		if (nBit & 1) == 1 {
			result = result * x
		}

		x = x * x
		nBit = nBit >> 1
	}

	if n < 0 {
		result = 1.0 / result
	}

	return result
}

func main() {
	fmt.Println(myPow(2.00000, 10))
	fmt.Println(myPow(2.10000, 3))
	fmt.Println(myPow(2.00000, -2))
}
