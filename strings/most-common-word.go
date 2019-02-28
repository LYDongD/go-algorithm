package main

import "fmt"

func mostCommonWord(paragraph string, banned []string) string {
	if len(paragraph) == 0 {
		return ""
	}

	wordCount := makeWordCount(paragraph)

	//deleteSensitiveWord
	for _, senseWord := range banned {
		if wordCount[senseWord] > 0 {
			delete(wordCount, senseWord)
		}
	}

	//max frequent
	max := 0
	result := ""
	for word, count := range wordCount {
		if count > max {
			max = count
			result = word
		}
	}

	return result
}

func makeWordCount(paragraph string) map[string]int {
	wordCount := make(map[string]int)
	var word []rune
	for _, char := range paragraph {
		if isLowerLetter(toLower(char)) {
			word = append(word, toLower(char))
		} else {
			if len(word) > 0 {
				wordCount[string(word)]++
			}
			word = []rune{}
		}
	}

	if len(word) > 0 {
		wordCount[string(word)]++
	}

	return wordCount
}

func isLowerLetter(char rune) bool {
	return char >= 'a' && char <= 'z'
}

func toLower(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return char + ('a' - 'A')
	}

	return char
}

func main() {
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
}
