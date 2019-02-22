package main

import "fmt"

func letterCasePermutation(S string) []string {
	if len(S) == 0 {
		return nil
	}

	var result []string
	return permutation(S, 0, "", result)
}

func permutation(S string, index int, currentStr string, result []string) []string {
	if index == len(S) {
		result = append(result, currentStr)
		return result
	}

	currentChar := rune(S[index])
	if !isCharacter(currentChar) {
		currentStr = currentStr + string(currentChar)
		result = permutation(S, index+1, currentStr, result)
	} else {
		result = permutation(S, index+1, currentStr+string(toLower(currentChar)), result)
		result = permutation(S, index+1, currentStr+string(toUpper(currentChar)), result)
	}

	return result
}

func isCharacter(char rune) bool {
	return (char >= 'a' && char < 'z') || (char >= 'A' && char <= 'Z')
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

	return char + ('a' - 'A')
}

func main() {
	fmt.Println(letterCasePermutation("a1b2"))
}
