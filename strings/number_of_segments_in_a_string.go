package strings

import (
	"unicode"
)

func CountSegments(s string) int {
	if len(s) == 0 {
		return 0
	}

	wordSpace := true
	segmentsCount := 0
	for _, char := range s {
		if unicode.IsSpace(char) {
			wordSpace = true
		} else {

			if wordSpace == true {
				segmentsCount++
			}
			wordSpace = false
		}
	}
	return segmentsCount
}
