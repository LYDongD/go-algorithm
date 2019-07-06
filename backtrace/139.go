package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	if len(wordDict) == 0 {
		return false
	}

	result := false
	wordHashDic := transferToHash(wordDict)
	return backtrace(s, wordHashDic, make(map[string]bool), result)
}

func transferToHash(wordDict []string) map[string]int {
	result := make(map[string]int)
	for _, word := range wordDict {
		result[word]++
	}

	return result
}

func backtrace(left string, wordHashDic map[string]int, cache map[string]bool, result bool) bool {
	if len(left) == 0 {
		return true
	}

	for i := 1; i <= len(left); i++ {
		if wordHashDic[left[:i]] == 0 {
			continue
		}

		cacheResult, ok := cache[left[i:]]
		if ok {
			result = result || cacheResult
		} else {
			current := backtrace(left[i:], wordHashDic, cache, result)
			cache[left[i:]] = current
			result = result || current
		}
		if result {
			break
		}
	}

	return result
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
