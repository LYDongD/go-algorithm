package main

import "fmt"

func minFlips(target string) int {
	if len(target) == 0 {
		return 0
	}

	count := 0
	for _, char := range target {
		if count%2 == int(char-'0') {
			continue
		}
		count++
	}

	return count
}

func main() {
	fmt.Println(minFlips("10111"))
	fmt.Println(minFlips("101"))
	fmt.Println(minFlips("00000"))
	fmt.Println(minFlips("001011101"))
}
