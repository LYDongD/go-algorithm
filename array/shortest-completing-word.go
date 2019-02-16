package main

import "fmt"

func shortestCompletingWord(licensePlate string, words []string) string {

	//1 licence预处理
	licenseDic := make(map[rune]int)
	for _, char := range licensePlate {
		if char >= 'A' && char <= 'Z' {
			licenseDic[char+('a'-'A')]++
		} else if char >= 'a' && char <= 'z' {
			licenseDic[char]++
		}
	}

	//2 线性遍历单词数组，构建单词统计字典，并选出第一个满足要求的最小长度的单词
	bestWord := ""
	for _, word := range words {
		if len(bestWord) > 0 && len(word) > len(bestWord) {
			continue
		}

		wordDic := makeWordDic(word)
		isQualified := true
		for char, count := range licenseDic {
			if wordDic[char] == 0 || wordDic[char] < count {
				isQualified = false
				break
			}
		}

		fmt.Println(word, isQualified)

		if isQualified && (len(bestWord) == 0 || len(word) < len(bestWord)) {
			bestWord = word
		}
	}

	return bestWord
}

func makeWordDic(word string) map[rune]int {
	dic := make(map[rune]int)
	for _, char := range word {
		dic[char]++
	}

	return dic
}

func main() {
	fmt.Println(shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"}))
	fmt.Println(shortestCompletingWord("1s3 456", []string{"looks", "pest", "stew", "show"}))
	fmt.Println(shortestCompletingWord("GrC8950", []string{"measure", "other", "every", "base", "according", "level", "meeting", "none", "marriage", "rest"}))
}
