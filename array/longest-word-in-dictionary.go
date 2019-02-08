package main

import "fmt"

func longestWord(words []string) string {
	if len(words) <= 1 {
		return ""
	}

	//建表
	dic := generateDic(words)

	result := ""
	for _, word := range words {
		//剪枝
		if len(word) < len(result) {
			continue
		}

		//查表
		if !isQualified(word, dic) {
			continue
		}

		//仅更新长的或排在前面的
		if len(word) > len(result) || word < result {
			result = word
		}
	}

	return result
}

func generateDic(words []string) map[string]bool {
	result := make(map[string]bool)
	for _, word := range words {
		result[word] = true
	}

	return result
}

func isQualified(word string, dic map[string]bool) bool {
	for i := 1; i < len(word); i++ {
		if !dic[word[:i]] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(longestWord([]string{"w", "wo", "wor", "worl", "world"}))
	fmt.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
}
