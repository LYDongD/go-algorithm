package main

import "fmt"

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	left := 0
	right := len(s) - 1
	for left < right {

		if !isAlphaNum(s[left]) {
			left++
		} else if !isAlphaNum(s[right]) {
			right--
		} else {
			if toLower(s[left]) != toLower(s[right]) {
				fmt.Println(s[left], s[right])
				return false
			} else {
				left++
				right--
			}
		}
	}

	return true
}

func toLower(a byte) int {
	runeInt := int(a)
	if runeInt >= 'A' && runeInt <= 'Z' {
		return runeInt + 'a' - 'A'
	}
	return runeInt
}

func isAlphaNum(char byte) bool {
	runeInt := int(char)
	if runeInt >= '0' && runeInt <= '9' {
		return true
	}

	if runeInt >= 'A' && runeInt <= 'Z' {
		return true
	}

	if runeInt >= 'a' && runeInt <= 'z' {
		return true
	}

	return false
}

func main() {
	fmt.Println(isPalindrome("OP"))
}
