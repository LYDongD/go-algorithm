package main

import "fmt"

func letterCombinations(digits string) []string {
    result := []string{}
    if len(digits) == 0 {
	return result
    }

    phoneDic := map[string]string{
	"2" : "abc",
	"3" : "def",
	"4" : "ghi",
	"5" : "jkl",
	"6" : "mno",
	"7" : "pqrs",
	"8" : "tuv",
	"9" : "wxyz",
    }

    return enumerateDigits(digits, phoneDic, "", result)
}


func enumerateDigits(digits string, phoneDic map[string]string, alphaStr string,  result []string) []string {
    if len(digits) == 0 {
	result = append(result, alphaStr)
	return result
    }

    alphas, ok := phoneDic[digits[:1]]
    if ok {
	for _, alpha := range alphas {
	    result = enumerateDigits(digits[1:], phoneDic, alphaStr + string(alpha), result)
	}
    }

    return result
}

func main() {
    fmt.Println(letterCombinations("23"))
}


