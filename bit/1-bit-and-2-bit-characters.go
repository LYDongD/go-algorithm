package main

import "fmt"

func isOneBitCharacter(bits []int) bool {
	if len(bits) == 0 {
		panic("bits array is empty")
	}

	if len(bits) == 1 {
		return true
	}

	//count 1 backwards
	oneCount := 0
	for i := len(bits) - 2; i >= 0; i-- {
		if bits[i] == 1 {
			oneCount++
		} else {
			break
		}
	}

	return oneCount%2 == 0
}

func main() {
	fmt.Println(isOneBitCharacter([]int{1, 0, 0}))
	fmt.Println(isOneBitCharacter([]int{1, 1, 1, 0}))
}
