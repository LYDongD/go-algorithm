package main

import "fmt"

func reorderSpaces(text string) string {
	if len(text) == 0 {
		return text
	}

	wordArr := []string{}
	spaceCount := 0
	wordCharCount := 0
	wordStart := -1
	for i, char := range text {
		if char == rune(' ') {
			spaceCount++
			if wordCharCount > 0 {
				wordArr = append(wordArr, text[wordStart:wordStart+wordCharCount])
				wordStart = -1
				wordCharCount = 0
			}
		} else {
			if wordCharCount == 0 {
				wordStart = i
			}
			wordCharCount++
		}
	}

	if wordCharCount > 0 {
		wordArr = append(wordArr, text[wordStart:wordStart+wordCharCount])
	}

	if spaceCount == 0 {
		return text
	}

	result := wordArr[0]
	spaceSplitCount := 0
	spaceLeftCount := 0
	if len(wordArr) > 1 {
		spaceSplitCount = spaceCount / (len(wordArr) - 1)
		spaceLeftCount = spaceCount % (len(wordArr) - 1)
	} else {
		spaceLeftCount = spaceCount
	}
	for i := 1; i < len(wordArr); i++ {
		for k := 0; k < spaceSplitCount; k++ {
			result = result + " "
		}

		result = result + wordArr[i]
	}

	for k := 0; k < spaceLeftCount; k++ {
		result = result + " "
	}

	return result
}

func main() {
	fmt.Println(reorderSpaces("  this   is  a sentence "))
	fmt.Println(reorderSpaces(" practice   makes   perfect"))
	fmt.Println(reorderSpaces("hello   world"))
	fmt.Println(reorderSpaces("  walks  udp package   into  bar a"))
	fmt.Println(reorderSpaces("a"))
}
