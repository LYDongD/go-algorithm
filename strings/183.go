package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	result := []string{}
	if len(s) <= 10 {
		return result
	}

	charDic := make(map[byte]int)
	charDic['A'] = 0
	charDic['C'] = 1
	charDic['G'] = 2
	charDic['T'] = 3

	cache := make(map[int]int)
	cursor := 0
	//skip first 9 char
	for i := 0; i < 9; i++ {
		cursor = cursor<<2 | charDic[s[i]]
	}

	//get last 10 chars
	for i := 9; i < len(s); i++ {
		cursor = cursor<<2 | charDic[s[i]]
		if cache[cursor&0xfffff] == 1 {
			result = append(result, s[i-10+1:i+1])
		}

		cache[cursor&0xfffff]++
	}

	return result
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAA"))
}
