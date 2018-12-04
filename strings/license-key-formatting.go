package main

import (
	"fmt"
	"strings"
)

func licenseKeyFormatting(S string, K int) string {
	//1 iterate from back to forward
	//2 count if not dash, and append
	//3 append dash when count == k
	//4 until all characters done
	result := ""
	count := 0
	for index := len(S) - 1; index >= 0; index-- {

		char := S[index]
		if char != '-' {
			if count == K {
				result = "-" + result
				count = 0
			}
			count++
			result = strings.ToUpper(string(char)) + result
		} else {
			continue
		}

	}

	return result
}

func main() {
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
	fmt.Println(licenseKeyFormatting("2-5g-3-J", 2))
	fmt.Println(licenseKeyFormatting("--a-a-a-a--", 2))
}
