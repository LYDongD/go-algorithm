package main

import "fmt"

func letterCasePermutation(S string) []string {
	result := []string{}
	chars := make([]rune, len(S))
	return permutation(S, 0, chars, result)
}

func permutation(S string, index int, chars []rune, result []string) []string {
	if index == len(S) {
	    return append(result, string(chars))
	}

	currentChar := rune(S[index])
	if !isCharacter(currentChar) {
		chars[index] = currentChar
		result = permutation(S, index+1, chars, result)
	}else {
	    chars[index] = toLower(currentChar)
	    result = permutation(S, index+1, chars, result)
	    chars[index] = toUpper(currentChar)
	    result = permutation(S, index+1, chars, result)
	}

	return result
}

func isCharacter(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func toLower(char rune) rune {
	if char >= 'a' && char <= 'z' {
		return char
	}

	return char + ('a' - 'A')
}

func toUpper(char rune) rune {
	if char >= 'A' && char <= 'Z' {
		return char
	}

	return char + ('A' - 'a')
}

func main() {
	fmt.Println(letterCasePermutation(""))
	fmt.Println(letterCasePermutation("a1b2"))
	fmt.Println(letterCasePermutation("3z4"))
	fmt.Println(letterCasePermutation("12345"))
}
