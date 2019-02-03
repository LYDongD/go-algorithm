package main

import "fmt"

func countBinarySubstrings(s string) int {
	if len(s) <= 0 {
		return 0
	}

	countZero := 0
	countOne := 0

	if s[0] == byte('0') {
		countZero++
	} else {
		countOne++
	}

	count := 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {

			count++

			//reset countZero or countOne
			if s[i] == byte('0') {
				countZero = 1
			} else {
				countOne = 1
			}
		} else {
			if s[i] == byte('0') {
				countZero++
				if countZero <= countOne {
					count++
				}
			} else {
				countOne++
				if countOne <= countZero {
					count++
				}
			}

		}
	}

	return count
}

func main() {
	fmt.Println(countBinarySubstrings("000111000"))
	fmt.Println(countBinarySubstrings("00110011"))
	fmt.Println(countBinarySubstrings("10101"))
	fmt.Println(countBinarySubstrings("11110011"))
	fmt.Println(countBinarySubstrings("100111001"))
}
