package main

import (
	"fmt"
)

func findComplement(num int) int {

	return num ^ mask(num)
}

//num's mask
func mask(num int) int {

	mask := 0
	for num > 0 {
		mask = mask<<1 + 1
		num = num >> 1
	}

	return mask
}

func main() {

	fmt.Println(mask(5), mask(1))

	fmt.Println(findComplement(5))
	fmt.Println(findComplement(1))

}
