package main

import (
	"fmt"
)

func findLUSlength(a string, b string) int {
	if len(a) != len(b) {
		if len(a) > len(b) {
			return len(a)
		} else {
			return len(b)
		}
	}

	for index := range a {
		if a[index] != b[index] {
			return len(a)
		}
	}

	return -1

}

func main() {
	fmt.Println(findLUSlength("aefawfawfawfaw", "aefawfeawfwafwaef"))
}
