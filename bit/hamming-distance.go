package main

import "fmt"

func hammingDistance(x int, y int) int {

	//xor
	t := x ^ y

	//count bit 1
	return oneBitCount(t)

}

func oneBitCount(t int) int {

	count := 0
	for t > 0 {
		if t&1 == 1 {
			count++
		}
		t = t >> 1
	}

	return count
}

func main() {

	fmt.Println(hammingDistance(1, 4))
	fmt.Println(hammingDistance(4, 4))

}
