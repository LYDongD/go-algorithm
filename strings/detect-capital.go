//package main
//
//import (
//	"fmt"
//)

func detectCapitalUse(word string) bool {
	if len(word) == 0 {
		return false
	}

	upperCaseCount := 0
	lowerCaseCount := 0
	firstRune := rune(0)
	for _, char := range word {

		if firstRune == 0 {
			firstRune = char
		}

		if isUpperCase(char) {
			upperCaseCount++
		} else {
			lowerCaseCount++
		}

	}

		if upperCaseCount >= 2 && lowerCaseCount >= 1 {
			return false
		}

		if !isUpperCase(firstRune) && upperCaseCount >= 1 {
			return false
		}

	return true
}

func isUpperCase(char rune) bool {
	if byte(char) <= byte(90) {
		return true
	} else {
		return false
	}
}

//func main() {
//	fmt.Println(detectCapitalUse("USA"))
//	fmt.Println(detectCapitalUse("FlaG"))
//	fmt.Println(detectCapitalUse("ffffF"))
//}
