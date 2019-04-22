package main

import "fmt"
import "math"

func bitwiseComplement(N int) int {
    if N <= 1 {
	return 1 - N
    }

    bitCount := countBit(N)
    return N ^ (int(math.Pow(2.0, float64(bitCount))) - 1)
}

func countBit(N int) int {
    count := 0
    for N > 0 {
	N = N >> 1
	count++
    }

    return count
}

func main() {
    fmt.Println(bitwiseComplement(5))
    fmt.Println(bitwiseComplement(7))
    fmt.Println(bitwiseComplement(10))
}
