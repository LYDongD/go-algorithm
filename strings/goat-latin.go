package main

import (
    "fmt"
    "strings"
)

func toGoatLatin(S string) string {
    if len(S) == 0 {
	return ""
    }

    words := strings.Split(S, " ")
    result  := ""
    for index, word := range words {
	firstChar := word[0]
	if !isVowel(firstChar) {
	    word = word[1:] + word[:1]
	}

	word = word + "ma"

	for i := 0; i < index + 1; i++ {
	    word = word + "a"
	}

	if index == 0 {
	    result = word
	}else {
	    result = result + " " + word
	}
    }

    return result
}

func isVowel(char byte) bool {
    if char >= 'A' && char <= 'Z' {
	char = char + ('a' - 'A')
    }

    return char == 'a' || char == 'e' || char == 'i'||
    char == 'o' || char == 'u'
}

func main() {
    fmt.Println(toGoatLatin("I speak Goat Latin"))
    fmt.Println(toGoatLatin("The quick brown fox jumped over the lazy dog"))
}


