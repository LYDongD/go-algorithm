package main

import "fmt"

func reverseOnlyLetters(S string) string {
    if len(S) <= 1 {
	return S
    }

    start := 0
    end := len(S) - 1
    s := []rune(S)
    for start < end {
	if isLetter(s[start]) && isLetter(s[end]) {
	    s[start], s[end] = s[end], s[start]
	    start++
	    end--
	}else {
	     if !isLetter(s[start]) {
	       start++
	     }

	     if !isLetter(s[end]) {
		end--
	     }
	}
    }

    return string(s)
}

func isLetter(char rune) bool {
    if char >= 'a' && char <= 'z' {
	return true
    }

    if char >= 'A' && char <= 'Z' {
	return true
    }

    return false
}

func main() {
    fmt.Println(reverseOnlyLetters("ab-cd"))
    fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj"))
    fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!"))
}
