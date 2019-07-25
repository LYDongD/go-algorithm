package main

import "fmt"

func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}

	wordList := extractWord(s)
	result := ""
	for i := len(wordList) - 1; i >= 0; i-- {
		result = result + wordList[i]
		if i != 0 {
			result = result + " "
		}
	}

	return result
}

func extractWord(s string) []string {
	wordList := []string{}
	wordStart, wordEnd := -1, -1
	for i, char := range s {
		if char != ' ' && wordStart == -1 {
			wordStart = i
		} else if char == ' ' && wordStart != -1 {
			wordEnd = i
			wordList = append(wordList, s[wordStart:wordEnd])
			wordStart = -1
		}
	}

	if wordStart != -1 {
		wordList = append(wordList, s[wordStart:])
	}

	return wordList
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("a good   example"))
	fmt.Println(reverseWords("  hello world!  "))
}
